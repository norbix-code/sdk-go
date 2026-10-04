package hub

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every Scheduler endpoint the gateway exposes, and the module method that
// calls it. Each case asserts the verb, the fully-resolved path (version
// substituted, route tokens filled), the query string and the JSON body, so a
// wrong path, a swapped verb or a parameter in the wrong place fails here
// instead of at runtime. Source of truth: gateway Hub.Scheduler/*.cs [Route].
//
// Nothing leaves the process: every call goes to a local test server.

const (
	testSchedulerTaskID = "tsk_1"
	testInitiatorUserID = "usr_1"
	testEmailTemplateID = "etpl_1"
)

type schedulerCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *SchedulerModule) error
}

// schedulerSaveRequest is a complete, valid save request: the initiator, the
// schedule and a typed email campaign task sent to all users.
func schedulerSaveRequest() map[string]any {
	return map[string]any{
		"initiatorUserId": testInitiatorUserID,
		"name":            "Weekly digest",
		"cron":            "0 9 * * 1",
		"isEnabled":       true,
		"stopOnError":     false,
		"task": dtos.EmailCampaignSchedulerTaskRequest{
			Type: dtos.SchedulerTaskTypeEmailCampaign,
			Campaign: &dtos.EmailCampaignRequest{
				Source:     dtos.EmailCampaignRecipientsSourceTypesAllUsers,
				TemplateId: testEmailTemplateID,
			},
		},
	}
}

func schedulerCases() []schedulerCase {
	return []schedulerCase{
		// --- module (PUT since the gateway moved them off GET)
		{"EnableScheduler", http.MethodPut, "/v2/scheduler/enable", nil, nil,
			func(ctx context.Context, m *SchedulerModule) error {
				return m.EnableScheduler(ctx, nil, nil)
			}},
		{"DisableScheduler", http.MethodPut, "/v2/scheduler/disable", nil, nil,
			func(ctx context.Context, m *SchedulerModule) error {
				return m.DisableScheduler(ctx, nil, nil)
			}},

		// --- tasks
		{"GetSchedulerTasks", http.MethodGet, "/v2/scheduler/tasks",
			url.Values{"pageSize": {"10"}, "type": {"EmailCampaign"}}, nil,
			func(ctx context.Context, m *SchedulerModule) error {
				return m.GetSchedulerTasks(ctx, map[string]any{"pageSize": 10, "type": "EmailCampaign"}, nil)
			}},
		{"GetSchedulerTask", http.MethodGet, "/v2/scheduler/tasks/" + testSchedulerTaskID, nil, nil,
			func(ctx context.Context, m *SchedulerModule) error {
				return m.GetSchedulerTask(ctx, testSchedulerTaskID, nil, nil)
			}},
		{"SaveSchedulerTask", http.MethodPost, "/v2/scheduler/tasks", nil,
			map[string]any{
				"initiatorUserId": testInitiatorUserID,
				"name":            "Weekly digest",
				"cron":            "0 9 * * 1",
				"isEnabled":       true,
				"task": map[string]any{
					"type": "EmailCampaign",
					"campaign": map[string]any{
						"source":     "AllUsers",
						"templateId": testEmailTemplateID,
					},
				},
			},
			func(ctx context.Context, m *SchedulerModule) error {
				return m.SaveSchedulerTask(ctx, schedulerSaveRequest(), nil)
			}},
		{"EnableSchedulerTask", http.MethodPut, "/v2/scheduler/tasks/" + testSchedulerTaskID + "/enable", nil, nil,
			func(ctx context.Context, m *SchedulerModule) error {
				return m.EnableSchedulerTask(ctx, testSchedulerTaskID, nil, nil)
			}},
		{"DisableSchedulerTask", http.MethodPut, "/v2/scheduler/tasks/" + testSchedulerTaskID + "/disable", nil, nil,
			func(ctx context.Context, m *SchedulerModule) error {
				return m.DisableSchedulerTask(ctx, testSchedulerTaskID, nil, nil)
			}},
		{"DeleteSchedulerTask", http.MethodDelete, "/v2/scheduler/tasks/" + testSchedulerTaskID, nil, nil,
			func(ctx context.Context, m *SchedulerModule) error {
				return m.DeleteSchedulerTask(ctx, testSchedulerTaskID, nil, nil)
			}},
	}
}

func newSchedulerTestModule(baseURL string) *SchedulerModule {
	return &SchedulerModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestSchedulerEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range schedulerCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newSchedulerTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}

			// Query: exactly the expected parameters, nothing else.
			gotQuery, _ := url.ParseQuery(gotRawQuery)
			if c.query == nil && gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if c.query != nil {
				gq, _ := json.Marshal(gotQuery)
				wq, _ := json.Marshal(c.query)
				if string(gq) != string(wq) {
					t.Errorf("query: got %s want %s", gq, wq)
				}
			}

			// Body: none for reads and toggles, the request for save.
			if c.body == nil {
				if len(gotBody) != 0 {
					t.Errorf("body: got %s want none", gotBody)
				}
				return
			}
			var got map[string]any
			if err := json.Unmarshal(gotBody, &got); err != nil {
				t.Fatalf("body is not JSON: %v (%s)", err, gotBody)
			}
			assertFields(t, got, c.body)
		})
	}
}

// The scheduler surface is 8 routes. If the gateway grows one and the module
// is regenerated, this count changes, so a new endpoint cannot arrive untested.
func TestSchedulerSurfaceSize(t *testing.T) {
	const want = 8
	if got := len(schedulerCases()); got != want {
		t.Errorf("scheduler endpoint count: got %d want %d", got, want)
	}
}

// The typed task reaches the wire with the discriminator the gateway reads
// ("type": "EmailCampaign", a string), the campaign nested under "campaign",
// and no field the caller did not set.
func TestSaveSchedulerTaskSendsTheTypedEmailTask(t *testing.T) {
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &got)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	if err := newSchedulerTestModule(srv.URL).SaveSchedulerTask(context.Background(), schedulerSaveRequest(), nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	task, _ := json.Marshal(got["task"])
	const want = `{"campaign":{"source":"AllUsers","templateId":"etpl_1"},"type":"EmailCampaign"}`
	if string(task) != want {
		t.Errorf("task: got %s want %s", task, want)
	}
	if _, ok := got["stopOnError"]; !ok {
		t.Errorf("stopOnError=false must still be sent: %v", got)
	}
}

// An update carries the task id in the body (there is no path token on save).
func TestSaveSchedulerTaskSendsTaskIDInTheBodyOnUpdate(t *testing.T) {
	var gotPath string
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &got)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	req := schedulerSaveRequest()
	req["taskId"] = testSchedulerTaskID
	if err := newSchedulerTestModule(srv.URL).SaveSchedulerTask(context.Background(), req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v2/scheduler/tasks" {
		t.Errorf("path: got %q", gotPath)
	}
	assertFields(t, got, map[string]any{"taskId": testSchedulerTaskID})
}
