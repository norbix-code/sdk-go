// Hand-written, not generated (gateway campaign audit/workflows).
//
// The two workflow (beta) helpers of Code.Workflows.md §15:
//
//   - (*CodeModule).WaitForWorkflowRun polls the run detail until the run is
//     finished or waiting. It is named WaitForWorkflowRun, not WaitForRun,
//     because the Code module also has function and marketplace runs.
//   - StepOutput[T] reads the output of a step from a run. Go methods cannot
//     take type parameters, so it is a package function, not a method.
//
// Workflows are in beta; bugs may occur.

package hub

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// Defaults of WaitForWorkflowRunOptions.
const (
	DefaultWorkflowRunWaitTimeout  = 60 * time.Second
	DefaultWorkflowRunPollInterval = time.Second
)

// WaitForWorkflowRunOptions tunes WaitForWorkflowRun. A zero value uses the
// defaults: a 60 s timeout and a poll every second.
type WaitForWorkflowRunOptions struct {
	// Timeout is how long to wait in total (default 60 s).
	Timeout time.Duration
	// PollInterval is the pause between two reads of the run (default 1 s).
	PollInterval time.Duration
}

// WorkflowRunTimeoutError is returned by WaitForWorkflowRun when the run is
// still running when the timeout ends. It carries the run as last read.
type WorkflowRunTimeoutError struct {
	Base *norbixerr.Error
	// RunID is the run that was waited on.
	RunID string
	// LastStatus is the run status of the last read (e.g. "running").
	LastStatus string
	// LastRun is the run as last read.
	LastRun *dtos.WorkflowRunDto
}

func (e *WorkflowRunTimeoutError) Error() string { return e.Base.Error() }

// Unwrap exposes the base *errors.Error (code NORBIX_WORKFLOW_RUN_TIMEOUT).
func (e *WorkflowRunTimeoutError) Unwrap() error { return e.Base }

// workflowRunSettled lists the run statuses WaitForWorkflowRun stops at: the
// three final ones and "waiting" (parked on a delay, an event or an approval —
// it will not move without a resume or the clock).
var workflowRunSettled = map[string]bool{
	"succeeded": true,
	"failed":    true,
	"cancelled": true,
	"waiting":   true,
}

// WaitForWorkflowRun polls GET /{version}/code/workflows/runs/{RunId} until
// the run's status is succeeded, failed, cancelled or waiting, and returns
// that run (steps included).
// Beta: workflows are in beta; bugs may occur.
//
// The first read is immediate. When the timeout ends first, the error is a
// *WorkflowRunTimeoutError carrying the last status and run. A refused read
// is returned as it came; a cancelled ctx returns ctx.Err(). An answer with
// no run is an *errors.Error with code NORBIX_WORKFLOW_RUN_MISSING.
func (m *CodeModule) WaitForWorkflowRun(ctx context.Context, runID string, opts WaitForWorkflowRunOptions) (*dtos.WorkflowRunDto, error) {
	if ctx == nil {
		ctx = context.Background()
	}
	if runID == "" {
		return nil, norbixerr.New("WaitForWorkflowRun needs a run id.", norbixerr.CodeValidation)
	}
	timeout := opts.Timeout
	if timeout <= 0 {
		timeout = DefaultWorkflowRunWaitTimeout
	}
	poll := opts.PollInterval
	if poll <= 0 {
		poll = DefaultWorkflowRunPollInterval
	}
	deadline := time.Now().Add(timeout)

	for {
		var resp dtos.WorkflowRunResponse
		if err := m.GetCodeWorkflowRun(ctx, runID, nil, &resp); err != nil {
			return nil, err
		}
		if resp.Run == nil {
			return nil, norbixerr.New(
				fmt.Sprintf("The answer for workflow run %s has no run.", runID),
				norbixerr.CodeWorkflowRunMissing,
			)
		}
		run := resp.Run
		if workflowRunSettled[run.Status] {
			return run, nil
		}

		remaining := time.Until(deadline)
		if remaining <= 0 {
			return nil, &WorkflowRunTimeoutError{
				Base: norbixerr.New(
					fmt.Sprintf("Workflow run %s is still %s after %s.", runID, run.Status, timeout),
					norbixerr.CodeWorkflowRunTimeout,
				),
				RunID:      runID,
				LastStatus: run.Status,
				LastRun:    run,
			}
		}
		// The last pause is cut to the deadline, so one more read happens at it.
		wait := poll
		if wait > remaining {
			wait = remaining
		}
		timer := time.NewTimer(wait)
		select {
		case <-ctx.Done():
			timer.Stop()
			return nil, ctx.Err()
		case <-timer.C:
		}
	}
}

// StepOutput decodes the output of step stepID from run into T.
// Beta: workflows are in beta; bugs may occur.
//
// A step runs once per attempt (retries) and once per pass (loops); the output
// read is the one of the latest execution with status "succeeded": the highest
// pass, then the highest attempt. ok is false (and value the zero T) when the
// step never succeeded in this run. The output is decoded by a JSON round trip,
// so T may be a struct with json tags, a map or any JSON-shaped type. An output
// cut by the gateway (WorkflowRunStepDto.OutputTruncated) is decoded as it is.
func StepOutput[T any](run *dtos.WorkflowRunDto, stepID string) (value T, ok bool, err error) {
	if run == nil {
		return value, false, nil
	}
	var best *dtos.WorkflowRunStepDto
	for _, step := range run.Steps {
		if step == nil || step.StepId != stepID || step.Status != "succeeded" {
			continue
		}
		if best == nil || step.Pass > best.Pass || (step.Pass == best.Pass && step.Attempt > best.Attempt) {
			best = step
		}
	}
	if best == nil {
		return value, false, nil
	}
	raw, err := json.Marshal(best.Output)
	if err != nil {
		return value, false, norbixerr.New("failed to encode the output of step "+stepID+": "+err.Error(), norbixerr.CodeError).WithWrapped(err)
	}
	if err := json.Unmarshal(raw, &value); err != nil {
		var zero T
		return zero, false, norbixerr.New("failed to decode the output of step "+stepID+": "+err.Error(), norbixerr.CodeError).WithWrapped(err)
	}
	return value, true, nil
}
