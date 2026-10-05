package api

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every Database endpoint the gateway exposes on the API (imports are out of
// scope), and the module method that calls it. Each case asserts the verb,
// the fully-resolved path (version substituted, route tokens filled), the
// query string, the JSON body, and the key + project headers, so a wrong
// path, a swapped verb or a parameter in the wrong place fails here instead
// of at runtime. Source of truth: the gateway endpoint manifest
// (sdks/typegen/coverage/endpoints.api.json, gateway branch audit/database).
//
// Nothing leaves the process: every call goes to a local test server.

type apiDatabaseCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *DatabaseModule) error
}

func apiDatabaseCases() []apiDatabaseCase {
	return []apiDatabaseCase{
		{"FindTerms", http.MethodGet, "/v2/database/taxonomies/categories/terms", url.Values{"pageSize": {"10"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindTerms(ctx, "categories", map[string]any{"pageSize": 10}, nil)
			}},
		{"FindTermsChildren", http.MethodGet, "/v2/database/taxonomies/categories/terms/t_1/children", url.Values{"pageSize": {"10"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindTermsChildren(ctx, "categories", "t_1", map[string]any{"pageSize": 10}, nil)
			}},
		{"FindTermTree", http.MethodGet, "/v2/database/taxonomies/categories/terms/tree", url.Values{"depth": {"2"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindTermTree(ctx, "categories", map[string]any{"depth": 2}, nil)
			}},
		{"FindTaxonomyTree", http.MethodGet, "/v2/database/taxonomies/tree", url.Values{"includeTerms": {"true"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindTaxonomyTree(ctx, map[string]any{"includeTerms": true}, nil)
			}},
		{"GetDatabaseSchema", http.MethodGet, "/v2/database/schemas/rec_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchema(ctx, "rec_1", nil, nil)
			}},
		{"GetDatabaseSchemas", http.MethodGet, "/v2/database/schemas", url.Values{"pageSize": {"10"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchemas(ctx, map[string]any{"pageSize": 10}, nil)
			}},
		{"Aggregate", http.MethodPost, "/v2/database/collections/products/aggregate", nil, map[string]any{"pipeline": "[]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.Aggregate(ctx, "products", map[string]any{"pipeline": "[]"}, nil)
			}},
		{"ChangeResponsibility", http.MethodPut, "/v2/database/collections/products/rec_1/responsibility", nil, map[string]any{"newResponsibleUserId": "usr_2"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ChangeResponsibility(ctx, "products", "rec_1", map[string]any{"newResponsibleUserId": "usr_2"}, nil)
			}},
		{"Count", http.MethodGet, "/v2/database/collections/products/count", url.Values{"filter": {"{}"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.Count(ctx, "products", map[string]any{"filter": "{}"}, nil)
			}},
		{"DeleteMany", http.MethodDelete, "/v2/database/collections/products/many", url.Values{"filter": {"{\"price\":10}"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteMany(ctx, "products", map[string]any{"filter": "{\"price\":10}"}, nil)
			}},
		{"DeleteOne", http.MethodDelete, "/v2/database/collections/products/rec_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteOne(ctx, "products", "rec_1", nil, nil)
			}},
		{"Distinct", http.MethodGet, "/v2/database/collections/products/distinct", url.Values{"field": {"brand"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.Distinct(ctx, "products", map[string]any{"field": "brand"}, nil)
			}},
		{"ExecuteAggregate", http.MethodPost, "/v2/database/collections/products/aggregates/agg_1/execute", nil, map[string]any{"tokens": map[string]any{"brand": "acme"}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ExecuteAggregate(ctx, "products", "agg_1", map[string]any{"tokens": map[string]any{"brand": "acme"}}, nil)
			}},
		{"Find", http.MethodGet, "/v2/database/collections/products", url.Values{"filter": {"{}"}, "pageSize": {"20"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.Find(ctx, "products", map[string]any{"filter": "{}", "pageSize": 20}, nil)
			}},
		{"FindOne", http.MethodGet, "/v2/database/collections/products/rec_1", url.Values{"databaseIntegrationId": {"int_1"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindOne(ctx, "products", "rec_1", map[string]any{"databaseIntegrationId": "int_1"}, nil)
			}},
		{"InsertMany", http.MethodPost, "/v2/database/collections/products/many", nil, map[string]any{"documents": "[{\"title\":\"A\"}]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.InsertMany(ctx, "products", map[string]any{"documents": "[{\"title\":\"A\"}]"}, nil)
			}},
		{"InsertOne", http.MethodPost, "/v2/database/collections/products", nil, map[string]any{"document": "{\"title\":\"Shoe\"}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.InsertOne(ctx, "products", map[string]any{"document": "{\"title\":\"Shoe\"}"}, nil)
			}},
		{"ReplaceOne", http.MethodPut, "/v2/database/collections/products/rec_1/replace", nil, map[string]any{"replacement": "{\"title\":\"Boot\"}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ReplaceOne(ctx, "products", "rec_1", map[string]any{"replacement": "{\"title\":\"Boot\"}"}, nil)
			}},
		{"UpdateMany", http.MethodPut, "/v2/database/collections/products/many", nil, map[string]any{"filter": "{}", "update": "{\"$set\":{\"price\":12}}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateMany(ctx, "products", map[string]any{"filter": "{}", "update": "{\"$set\":{\"price\":12}}"}, nil)
			}},
		{"UpdateOne", http.MethodPut, "/v2/database/collections/products/rec_1", nil, map[string]any{"update": "{\"$set\":{\"price\":12}}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateOne(ctx, "products", "rec_1", map[string]any{"update": "{\"$set\":{\"price\":12}}"}, nil)
			}},
		{"FindOwn", http.MethodGet, "/v2/database/collections/products/own", url.Values{"filter": {"{}"}, "pageSize": {"20"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindOwn(ctx, "products", map[string]any{"filter": "{}", "pageSize": 20}, nil)
			}},
		{"FindMergedTermTree", http.MethodGet, "/v2/database/taxonomies/categories/merged-tree", url.Values{"databaseIntegrationId": {"int_1"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindMergedTermTree(ctx, "categories", map[string]any{"databaseIntegrationId": "int_1"}, nil)
			}},
	}
}

func newAPIDatabaseTestModule(baseURL string) *DatabaseModule {
	return &DatabaseModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: baseURL,
		BaseURLHub: "http://hub.invalid",
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestAPIDatabaseEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range apiDatabaseCases() {
		t.Run(c.name, func(t *testing.T) {
			var gotMethod, gotPath, gotRawQuery, gotAuth, gotProject string
			var gotBody []byte
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotMethod = r.Method
				gotPath = r.URL.Path
				gotRawQuery = r.URL.RawQuery
				gotAuth = r.Header.Get("Authorization")
				gotProject = r.Header.Get("X-CM-ProjectId")
				gotBody, _ = io.ReadAll(r.Body)
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()

			if err := c.call(context.Background(), newAPIDatabaseTestModule(srv.URL)); err != nil {
				t.Fatalf("%s: unexpected error: %v", c.name, err)
			}
			if gotMethod != c.verb {
				t.Errorf("verb: got %q want %q", gotMethod, c.verb)
			}
			if gotPath != c.path {
				t.Errorf("path: got %q want %q", gotPath, c.path)
			}
			if gotAuth != "Bearer key_1" {
				t.Errorf("authorization: got %q want %q", gotAuth, "Bearer key_1")
			}
			if gotProject != "proj_1" {
				t.Errorf("project header: got %q want %q", gotProject, "proj_1")
			}

			// Query: exactly the expected parameters, nothing else.
			if c.query == nil && gotRawQuery != "" {
				t.Errorf("query: got %q want none", gotRawQuery)
			}
			if c.query != nil {
				gotQuery, _ := url.ParseQuery(gotRawQuery)
				gq, _ := json.Marshal(gotQuery)
				wq, _ := json.Marshal(c.query)
				if string(gq) != string(wq) {
					t.Errorf("query: got %s want %s", gq, wq)
				}
			}

			// Body: none for reads, deletes and bare toggles; the request otherwise.
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
			assertDatabaseFields(t, got, c.body)
		})
	}
}

// The API database surface is 22 routes (imports excluded). If the gateway
// grows one and a method is added, this count changes, so a new endpoint
// cannot arrive untested.
func TestAPIDatabaseSurfaceSize(t *testing.T) {
	const want = 22
	if got := len(apiDatabaseCases()); got != want {
		t.Errorf("database endpoint count: got %d want %d", got, want)
	}
}

// assertDatabaseFields fails for every key of want that is missing from got or
// holds a different value. Values are compared as JSON.
func assertDatabaseFields(t *testing.T, got, want map[string]any) {
	t.Helper()
	for key, w := range want {
		g, ok := got[key]
		if !ok {
			t.Errorf("field %q missing from body", key)
			continue
		}
		gb, _ := json.Marshal(g)
		wb, _ := json.Marshal(w)
		if string(gb) != string(wb) {
			t.Errorf("field %q: got %s want %s", key, gb, wb)
		}
	}
}
