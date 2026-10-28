package dtos

import (
	"encoding/json"
	"testing"
)

// A workflow step's input and output are any JSON value: a transform or
// condition step stores its expression string as input, and a transform may
// return a number. v2.41.0 typed them as a JSON object and refused the run.
func TestWorkflowRunReadsNonObjectStepValues(t *testing.T) {
	body := `{"run":{"runId":"run_1","status":"succeeded","output":42,` +
		`"request":{"amount":21},"trigger":{"kind":"manual"},"steps":[` +
		`{"stepId":"total","kind":"transform","status":"succeeded","input":"$.amount * 2","output":42}]}}`

	var res WorkflowRunResponse
	if err := json.Unmarshal([]byte(body), &res); err != nil {
		t.Fatalf("unmarshal: %v", err)
	}
	step := res.Run.Steps[0]
	if got, ok := step.Input.(string); !ok || got != "$.amount * 2" {
		t.Errorf("step input: got %#v", step.Input)
	}
	if got, ok := step.Output.(float64); !ok || got != 42 {
		t.Errorf("step output: got %#v", step.Output)
	}
	if got, ok := res.Run.Output.(float64); !ok || got != 42 {
		t.Errorf("run output: got %#v", res.Run.Output)
	}
	if res.Run.Request["amount"] != float64(21) {
		t.Errorf("request: got %#v", res.Run.Request)
	}
}
