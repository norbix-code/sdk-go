package hub

import (
	"context"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"sync/atomic"
	"testing"
	"time"

	norbixerr "github.com/norbix-code/sdk-go/v2/norbix/errors"
	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// The workflow (beta) helpers of Code.Workflows.md §15: WaitForWorkflowRun
// polls the run detail until the run is finished or waiting; StepOutput reads
// the output of a step's latest successful execution. Nothing leaves the
// process: every call goes to a local test server, with millisecond polls.

// runServer answers GET /v2/code/workflows/runs/{runId} with the bodies in
// order (the last one repeats) and counts the reads.
func runServer(t *testing.T, bodies ...string) (*httptest.Server, *int32) {
	t.Helper()
	var hits int32
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet || r.URL.Path != "/v2/code/workflows/runs/"+testWorkflowRunID {
			t.Errorf("unexpected request %s %s", r.Method, r.URL.Path)
		}
		n := int(atomic.AddInt32(&hits, 1))
		if n > len(bodies) {
			n = len(bodies)
		}
		w.Header().Set("Content-Type", "application/json")
		_, _ = io.WriteString(w, bodies[n-1])
	}))
	t.Cleanup(srv.Close)
	return srv, &hits
}

const runningBody = `{"run":{"runId":"` + testWorkflowRunID + `","status":"running","steps":[{"stepId":"charge","status":"running","attempt":1,"pass":1}]}}`

func TestWaitForWorkflowRunPollsUntilSucceeded(t *testing.T) {
	succeeded := `{"run":{"runId":"` + testWorkflowRunID + `","status":"succeeded","steps":[` +
		`{"stepId":"charge","status":"failed","attempt":1,"pass":1},` +
		`{"stepId":"charge","status":"succeeded","attempt":2,"pass":1,"output":{"paymentId":"pay_1","amountMinor":4200}}]}}`
	srv, hits := runServer(t, runningBody, runningBody, succeeded)

	run, err := newCodeTestModule(srv.URL).WaitForWorkflowRun(context.Background(), testWorkflowRunID,
		WaitForWorkflowRunOptions{Timeout: 5 * time.Second, PollInterval: time.Millisecond})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if run.Status != "succeeded" || run.RunId != testWorkflowRunID {
		t.Errorf("run: got %s %s", run.RunId, run.Status)
	}
	if got := atomic.LoadInt32(hits); got != 3 {
		t.Errorf("reads: got %d want 3", got)
	}

	type charge struct {
		PaymentID   string `json:"paymentId"`
		AmountMinor int    `json:"amountMinor"`
	}
	out, ok, err := StepOutput[charge](run, "charge")
	if err != nil || !ok {
		t.Fatalf("StepOutput: ok=%v err=%v", ok, err)
	}
	if out.PaymentID != "pay_1" || out.AmountMinor != 4200 {
		t.Errorf("output: got %+v", out)
	}
}

func TestWaitForWorkflowRunStopsAtWaiting(t *testing.T) {
	waiting := `{"run":{"runId":"` + testWorkflowRunID + `","status":"waiting","waitingOn":[{"stepId":"approve","kind":"approval","assignees":["usr_1"]}]}}`
	srv, hits := runServer(t, waiting, runningBody)

	run, err := newCodeTestModule(srv.URL).WaitForWorkflowRun(context.Background(), testWorkflowRunID,
		WaitForWorkflowRunOptions{Timeout: time.Second, PollInterval: time.Millisecond})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if run.Status != "waiting" || len(run.WaitingOn) != 1 || run.WaitingOn[0].Kind != "approval" {
		t.Errorf("run: got %+v", run)
	}
	if got := atomic.LoadInt32(hits); got != 1 {
		t.Errorf("reads: got %d want 1", got)
	}
}

func TestWaitForWorkflowRunTimesOut(t *testing.T) {
	srv, hits := runServer(t, runningBody)

	start := time.Now()
	run, err := newCodeTestModule(srv.URL).WaitForWorkflowRun(context.Background(), testWorkflowRunID,
		WaitForWorkflowRunOptions{Timeout: 30 * time.Millisecond, PollInterval: 2 * time.Millisecond})
	if run != nil {
		t.Errorf("run: got %+v want nil", run)
	}
	var timeoutErr *WorkflowRunTimeoutError
	if !errors.As(err, &timeoutErr) {
		t.Fatalf("error: got %T %v want *WorkflowRunTimeoutError", err, err)
	}
	if timeoutErr.RunID != testWorkflowRunID || timeoutErr.LastStatus != "running" {
		t.Errorf("timeout error: got run %q status %q", timeoutErr.RunID, timeoutErr.LastStatus)
	}
	if timeoutErr.LastRun == nil || timeoutErr.LastRun.Status != "running" {
		t.Errorf("last run: got %+v", timeoutErr.LastRun)
	}
	var base *norbixerr.Error
	if !errors.As(err, &base) || base.Code != norbixerr.CodeWorkflowRunTimeout {
		t.Errorf("base error: got %+v", base)
	}
	if got := atomic.LoadInt32(hits); got < 2 {
		t.Errorf("reads: got %d want several", got)
	}
	if elapsed := time.Since(start); elapsed > 2*time.Second {
		t.Errorf("took %s", elapsed)
	}
}

func TestWaitForWorkflowRunRespectsContext(t *testing.T) {
	srv, _ := runServer(t, runningBody)
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Millisecond)
	defer cancel()

	_, err := newCodeTestModule(srv.URL).WaitForWorkflowRun(ctx, testWorkflowRunID,
		WaitForWorkflowRunOptions{Timeout: time.Minute, PollInterval: 5 * time.Millisecond})
	if err == nil {
		t.Fatal("expected an error")
	}
	var timeoutErr *WorkflowRunTimeoutError
	if errors.As(err, &timeoutErr) {
		t.Errorf("a cancelled ctx is not the helper's timeout: %v", err)
	}
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("error: got %v want context.DeadlineExceeded", err)
	}
}

func TestWaitForWorkflowRunWithoutRunIsAnError(t *testing.T) {
	srv, _ := runServer(t, `{}`)
	_, err := newCodeTestModule(srv.URL).WaitForWorkflowRun(context.Background(), testWorkflowRunID,
		WaitForWorkflowRunOptions{Timeout: time.Second, PollInterval: time.Millisecond})
	var base *norbixerr.Error
	if !errors.As(err, &base) || base.Code != norbixerr.CodeWorkflowRunMissing {
		t.Errorf("error: got %v want code %s", err, norbixerr.CodeWorkflowRunMissing)
	}
}

func TestStepOutputPicksTheLatestSuccessfulExecution(t *testing.T) {
	step := func(status string, pass, attempt float64, n int) *dtos.WorkflowRunStepDto {
		return &dtos.WorkflowRunStepDto{StepId: "loop", Status: status, Pass: pass, Attempt: attempt, Output: map[string]any{"n": n}}
	}
	cases := []struct {
		name   string
		steps  []*dtos.WorkflowRunStepDto
		wantN  int
		wantOK bool
	}{
		{"highest pass wins over a higher attempt",
			[]*dtos.WorkflowRunStepDto{step("succeeded", 1, 3, 13), step("succeeded", 2, 1, 21), step("succeeded", 1, 1, 11)}, 21, true},
		{"highest attempt within the highest pass",
			[]*dtos.WorkflowRunStepDto{step("succeeded", 2, 1, 21), step("succeeded", 2, 2, 22), step("succeeded", 1, 5, 15)}, 22, true},
		{"failed executions are skipped",
			[]*dtos.WorkflowRunStepDto{step("failed", 3, 1, 31), step("succeeded", 2, 2, 22), step("cancelled", 2, 3, 23)}, 22, true},
		{"never succeeded",
			[]*dtos.WorkflowRunStepDto{step("failed", 1, 1, 11), step("running", 1, 2, 12)}, 0, false},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			other := &dtos.WorkflowRunStepDto{StepId: "other", Status: "succeeded", Pass: 9, Attempt: 9, Output: map[string]any{"n": 99}}
			run := &dtos.WorkflowRunDto{Steps: append(c.steps, other)}
			got, ok, err := StepOutput[map[string]int](run, "loop")
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if ok != c.wantOK {
				t.Fatalf("ok: got %v want %v", ok, c.wantOK)
			}
			if !ok {
				if got != nil {
					t.Errorf("value: got %v want the zero value", got)
				}
				return
			}
			if got["n"] != c.wantN {
				t.Errorf("n: got %d want %d", got["n"], c.wantN)
			}
		})
	}
}

func TestStepOutputUnknownStepAndNilRun(t *testing.T) {
	if _, ok, err := StepOutput[map[string]any](nil, "x"); ok || err != nil {
		t.Errorf("nil run: ok=%v err=%v", ok, err)
	}
	run := &dtos.WorkflowRunDto{Steps: []*dtos.WorkflowRunStepDto{{StepId: "a", Status: "succeeded"}}}
	if _, ok, err := StepOutput[map[string]any](run, "b"); ok || err != nil {
		t.Errorf("unknown step: ok=%v err=%v", ok, err)
	}
}

func TestStepOutputDecodeError(t *testing.T) {
	run := &dtos.WorkflowRunDto{Steps: []*dtos.WorkflowRunStepDto{{StepId: "a", Status: "succeeded", Output: map[string]any{"n": "text"}}}}
	type typed struct {
		N int `json:"n"`
	}
	got, ok, err := StepOutput[typed](run, "a")
	if err == nil || ok || got != (typed{}) {
		t.Errorf("got %+v ok=%v err=%v want a decode error", got, ok, err)
	}
}
