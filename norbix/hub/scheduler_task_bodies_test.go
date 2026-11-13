package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
)

// The "Code call" and "Workflow call" task bodies, built from the generated
// dtos, reach the wire under "task" with the discriminator the gateway's
// SchedulerTaskRequestJsonConverter reads (CodeFunctionalCall, WorkflowCall)
// and the fields of Hub.Scheduler/Tasks/Save.Code.cs and Save.Workflow.cs.
// Twins: sdk-ts tests/hub/scheduler-task-bodies.test.ts, sdk-python
// test_save_scheduler_task_sends_code_and_workflow_tasks.
func TestSaveSchedulerTaskSendsCodeAndWorkflowTasks(t *testing.T) {
	cases := []struct {
		name string
		task any
		want string
	}{
		{"CodeFunctionalCall", dtos.CodeFunctionSchedulerTaskRequest{
			Type:           dtos.SchedulerTaskTypeCodeFunctionalCall,
			FunctionRef:    &dtos.FunctionRefDto{IntegrationId: "int_lambda", FunctionId: "nightly-report", Qualifier: "live"},
			InvocationType: dtos.CodeInvocationTypeAsync,
			Config:         `{"source":"go"}`,
		}, `{"config":"{\"source\":\"go\"}","functionRef":{"functionId":"nightly-report","integrationId":"int_lambda","qualifier":"live"},"invocationType":"Async","type":"CodeFunctionalCall"}`},
		{"WorkflowCall", dtos.WorkflowSchedulerTaskRequest{
			Type:        dtos.SchedulerTaskTypeWorkflowCall,
			WorkflowRef: "wf_3kQ9mZ2vXc8LpR4wT7yH1n",
			Payload:     map[string]string{"source": "go"},
		}, `{"payload":{"source":"go"},"type":"WorkflowCall","workflowRef":"wf_3kQ9mZ2vXc8LpR4wT7yH1n"}`},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotQuery string
			var got map[string]any
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod, gotPath, gotQuery = r.Method, r.URL.Path, r.URL.RawQuery
				b, _ := io.ReadAll(r.Body)
				_ = json.Unmarshal(b, &got)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			req := schedulerSaveRequest()
			req["task"] = c.task
			if err := newSchedulerTestModule(srv.URL).SaveSchedulerTask(context.Background(), req, nil); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotMethod != http.MethodPost || gotPath != "/v2/scheduler/tasks" || gotQuery != "" {
				t.Errorf("route: got %s %s?%s", gotMethod, gotPath, gotQuery)
			}
			task, _ := json.Marshal(got["task"])
			if string(task) != c.want {
				t.Errorf("task: got %s want %s", task, c.want)
			}
			assertFields(t, got, map[string]any{"name": "Weekly digest", "cron": "0 9 * * 1", "initiatorUserId": testInitiatorUserID})
		})
	}
}
