package hub

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"
	"time"

	nberrors "github.com/norbix-code/sdk-go/v2/norbix/errors"
	"github.com/norbix-code/sdk-go/v2/norbix/hub/dtos"
	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// Every Database endpoint the gateway exposes on the HUB (imports are out of
// scope), and the module method that calls it. Each case asserts the verb,
// the fully-resolved path (version substituted, route tokens filled), the
// query string, the JSON body, and the key + project headers, so a wrong
// path, a swapped verb or a parameter in the wrong place fails here instead
// of at runtime. Source of truth: the gateway endpoint manifest
// (sdks/typegen/coverage/endpoints.hub.json, gateway branch audit/database).
//
// Nothing leaves the process: every call goes to a local test server.

type hubDatabaseCase struct {
	name  string
	verb  string
	path  string
	query url.Values     // nil: no query string expected
	body  map[string]any // nil: no body expected
	call  func(ctx context.Context, m *DatabaseModule) error
}

func hubDatabaseCases() []hubDatabaseCase {
	return []hubDatabaseCase{
		{"DisableDatabase", http.MethodPut, "/v2/database/disable", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DisableDatabase(ctx, nil, nil)
			}},
		{"EnableDatabase", http.MethodPut, "/v2/database/enable", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.EnableDatabase(ctx, nil, nil)
			}},
		{"DeleteSchemaTrigger", http.MethodDelete, "/v2/database/schemas/triggers/trg_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteSchemaTrigger(ctx, "trg_1", nil, nil)
			}},
		{"DisableSchemaTrigger", http.MethodPatch, "/v2/database/schemas/triggers/trg_1/disable", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DisableSchemaTrigger(ctx, "trg_1", nil, nil)
			}},
		{"EnableSchemaTrigger", http.MethodPatch, "/v2/database/schemas/triggers/trg_1/enable", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.EnableSchemaTrigger(ctx, "trg_1", nil, nil)
			}},
		{"GetSchemaTrigger", http.MethodGet, "/v2/database/schemas/triggers/rec_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetSchemaTrigger(ctx, "rec_1", nil, nil)
			}},
		{"GetSchemaTriggers", http.MethodGet, "/v2/database/schemas/triggers", url.Values{"schemaId": {"sch_1"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetSchemaTriggers(ctx, map[string]any{"schemaId": "sch_1"}, nil)
			}},
		{"SaveSchemaTrigger", http.MethodPost, "/v2/database/schemas/triggers", nil, map[string]any{"trigger": map[string]any{"schemaId": "sch_1", "name": "On insert"}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SaveSchemaTrigger(ctx, map[string]any{"trigger": map[string]any{"schemaId": "sch_1", "name": "On insert"}}, nil)
			}},
		{"DeleteDatabaseTaxonomy", http.MethodDelete, "/v2/database/taxonomies/id_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteDatabaseTaxonomy(ctx, "id_1", nil, nil)
			}},
		{"GetDatabaseTaxonomy", http.MethodGet, "/v2/database/taxonomies/rec_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseTaxonomy(ctx, "rec_1", nil, nil)
			}},
		{"GetDatabaseTaxonomies", http.MethodGet, "/v2/database/taxonomies", url.Values{"pageSize": {"10"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseTaxonomies(ctx, map[string]any{"pageSize": 10}, nil)
			}},
		{"SaveDatabaseTaxonomy", http.MethodPost, "/v2/database/taxonomies", nil, map[string]any{"taxonomyName": "categories", "description": "Product categories", "dependencies": []any{"brands"}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SaveDatabaseTaxonomy(ctx, map[string]any{"taxonomyName": "categories", "description": "Product categories", "dependencies": []any{"brands"}}, nil)
			}},
		{"DeleteDatabaseTaxonomyTerm", http.MethodDelete, "/v2/database/taxonomies/tax_1/terms/id_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteDatabaseTaxonomyTerm(ctx, "tax_1", "id_1", nil, nil)
			}},
		{"DeleteManyDatabaseTaxonomyTerms", http.MethodDelete, "/v2/database/taxonomies/tax_1/terms/many", url.Values{"filter": {"{\"name\":\"old\"}"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteManyDatabaseTaxonomyTerms(ctx, "tax_1", map[string]any{"filter": "{\"name\":\"old\"}"}, nil)
			}},
		{"GetDatabaseTaxonomyTerm", http.MethodGet, "/v2/database/taxonomies/tax_1/terms/id_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseTaxonomyTerm(ctx, "tax_1", "id_1", nil, nil)
			}},
		{"SaveDatabaseTaxonomyTerm", http.MethodPost, "/v2/database/taxonomies/tax_1/terms", nil, map[string]any{"document": "{\"name\":\"shoes\"}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SaveDatabaseTaxonomyTerm(ctx, "tax_1", map[string]any{"document": "{\"name\":\"shoes\"}"}, nil)
			}},
		{"UpdateDatabaseTaxonomyTerm", http.MethodPut, "/v2/database/taxonomies/tax_1/terms/id_1", nil, map[string]any{"update": "{\"$set\":{\"name\":\"boots\"}}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateDatabaseTaxonomyTerm(ctx, "tax_1", "id_1", map[string]any{"update": "{\"$set\":{\"name\":\"boots\"}}"}, nil)
			}},
		{"DeleteDatabaseSchema", http.MethodDelete, "/v2/database/schemas/id_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteDatabaseSchema(ctx, "id_1", nil, nil)
			}},
		{"DiscardDatabaseSchemaDraft", http.MethodDelete, "/v2/database/schemas/id_1/draft", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DiscardDatabaseSchemaDraft(ctx, "id_1", nil, nil)
			}},
		{"GetDatabaseSchema", http.MethodGet, "/v2/database/schemas/rec_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchema(ctx, "rec_1", nil, nil)
			}},
		{"GetDatabaseSchemas", http.MethodGet, "/v2/database/schemas", url.Values{"pageSize": {"10"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchemas(ctx, map[string]any{"pageSize": 10}, nil)
			}},
		{"GetDatabaseSchemaDraft", http.MethodGet, "/v2/database/schemas/id_1/draft", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchemaDraft(ctx, "id_1", nil, nil)
			}},
		{"GetDatabaseSchemaVersionDiff", http.MethodGet, "/v2/database/schemas/id_1/versions/diff", url.Values{"fromVersion": {"1"}, "toVersion": {"2"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchemaVersionDiff(ctx, "id_1", map[string]any{"fromVersion": 1, "toVersion": 2}, nil)
			}},
		{"GetDatabaseSchemaVersions", http.MethodGet, "/v2/database/schemas/id_1/versions", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchemaVersions(ctx, "id_1", nil, nil)
			}},
		{"PublishDatabaseSchema", http.MethodPost, "/v2/database/schemas/id_1/publish", nil, map[string]any{"confirmed": true},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.PublishDatabaseSchema(ctx, "id_1", map[string]any{"confirmed": true}, nil)
			}},
		{"RenameDatabaseSchema", http.MethodPut, "/v2/database/schemas/id_1/rename", nil, map[string]any{"title": "Catalog"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.RenameDatabaseSchema(ctx, "id_1", map[string]any{"title": "Catalog"}, nil)
			}},
		{"SaveDatabaseSchema", http.MethodPost, "/v2/database/schemas", nil, map[string]any{"schemaName": "products", "dataSchema": "{\"type\":\"object\"}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SaveDatabaseSchema(ctx, map[string]any{"schemaName": "products", "dataSchema": "{\"type\":\"object\"}"}, nil)
			}},
		{"UpdateDatabaseSchemaDraft", http.MethodPut, "/v2/database/schemas/id_1/draft", nil, map[string]any{"dataSchema": "{\"type\":\"object\"}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateDatabaseSchemaDraft(ctx, "id_1", map[string]any{"dataSchema": "{\"type\":\"object\"}"}, nil)
			}},
		{"UpdateDatabaseSchemaSettings", http.MethodPut, "/v2/database/schemas/id_1/settings", nil, map[string]any{"settings": map[string]any{"translatable": true}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateDatabaseSchemaSettings(ctx, "id_1", map[string]any{"settings": map[string]any{"translatable": true}}, nil)
			}},
		{"DeleteDatabaseIntegration", http.MethodDelete, "/v2/database/integrations/id_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteDatabaseIntegration(ctx, "id_1", nil, nil)
			}},
		{"DisableDatabaseIntegration", http.MethodPut, "/v2/database/integrations/id_1/disable", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DisableDatabaseIntegration(ctx, "id_1", nil, nil)
			}},
		{"EnableDatabaseIntegration", http.MethodPut, "/v2/database/integrations/id_1/enable", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.EnableDatabaseIntegration(ctx, "id_1", nil, nil)
			}},
		{"GetDatabaseIntegration", http.MethodGet, "/v2/database/integrations/rec_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseIntegration(ctx, "rec_1", nil, nil)
			}},
		{"GetDatabaseIntegrations", http.MethodGet, "/v2/database/integrations", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseIntegrations(ctx, nil, nil)
			}},
		{"SaveDatabaseIntegration", http.MethodPost, "/v2/database/integrations", nil, map[string]any{"integration": map[string]any{"name": "Main"}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SaveDatabaseIntegration(ctx, map[string]any{"integration": map[string]any{"name": "Main"}}, nil)
			}},
		{"SetDatabaseIntegrationAsDefault", http.MethodPut, "/v2/database/integrations/id_1/default", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SetDatabaseIntegrationAsDefault(ctx, "id_1", nil, nil)
			}},
		{"DeleteDatabaseAggregate", http.MethodDelete, "/v2/database/aggregates/id_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteDatabaseAggregate(ctx, "id_1", nil, nil)
			}},
		{"GetDatabaseAggregate", http.MethodGet, "/v2/database/aggregates/id_1", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseAggregate(ctx, "id_1", nil, nil)
			}},
		{"GetDatabaseAggregates", http.MethodGet, "/v2/database/aggregates", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseAggregates(ctx, nil, nil)
			}},
		{"SaveDatabaseAggregate", http.MethodPost, "/v2/database/aggregates", nil, map[string]any{"displayName": "Top products", "schemaId": "sch_1", "pipeline": "[]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SaveDatabaseAggregate(ctx, map[string]any{"displayName": "Top products", "schemaId": "sch_1", "pipeline": "[]"}, nil)
			}},
		{"TestDatabaseAggregate", http.MethodPost, "/v2/database/aggregates/test", nil, map[string]any{"collectionName": "products", "pipeline": "[]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.TestDatabaseAggregate(ctx, map[string]any{"collectionName": "products", "pipeline": "[]"}, nil)
			}},
		{"GetAllowedFlexTiers", http.MethodGet, "/v2/database/integrations/flex-tiers", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetAllowedFlexTiers(ctx, nil, nil)
			}},
		{"TestDatabaseIntegration", http.MethodPost, "/v2/database/integrations/test", nil, map[string]any{"integrationId": "int_1"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.TestDatabaseIntegration(ctx, map[string]any{"integrationId": "int_1"}, nil)
			}},
		{"RevealManagedFlexConnectionString", http.MethodGet, "/v2/database/integrations/id_1/connection-string", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.RevealManagedFlexConnectionString(ctx, "id_1", nil, nil)
			}},
		{"SeedCollectionRecords", http.MethodPost, "/v2/database/collections/seed", nil, map[string]any{"mode": "Sample", "collections": "products"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.SeedCollectionRecords(ctx, map[string]any{"mode": "Sample", "collections": "products"}, nil)
			}},
		{"FindRecords", http.MethodGet, "/v2/database/collections/products", url.Values{"filter": {"{\"price\":{\"$gt\":10}}"}, "pageSize": {"20"}, "sortBy": {"price"}, "expandReferences": {"true"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindRecords(ctx, "products", map[string]any{"filter": "{\"price\":{\"$gt\":10}}", "pageSize": 20, "sortBy": "price", "expandReferences": true}, nil)
			}},
		{"InsertRecord", http.MethodPost, "/v2/database/collections/products", nil, map[string]any{"document": "{\"title\":\"Shoe\"}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.InsertRecord(ctx, "products", map[string]any{"document": "{\"title\":\"Shoe\"}"}, nil)
			}},
		{"AggregateRecords", http.MethodPost, "/v2/database/collections/products/aggregate", nil, map[string]any{"pipeline": "[{\"$match\":{}}]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.AggregateRecords(ctx, "products", map[string]any{"pipeline": "[{\"$match\":{}}]"}, nil)
			}},
		{"ExecuteRecordsAggregate", http.MethodPost, "/v2/database/collections/products/aggregates/agg_1/execute", nil, map[string]any{"tokens": map[string]any{"brand": "acme"}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ExecuteRecordsAggregate(ctx, "products", "agg_1", map[string]any{"tokens": map[string]any{"brand": "acme"}}, nil)
			}},
		{"CountRecords", http.MethodGet, "/v2/database/collections/products/count", url.Values{"filter": {"{\"price\":10}"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.CountRecords(ctx, "products", map[string]any{"filter": "{\"price\":10}"}, nil)
			}},
		{"DistinctRecordValues", http.MethodGet, "/v2/database/collections/products/distinct", url.Values{"field": {"brand"}, "filter": {"{}"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DistinctRecordValues(ctx, "products", map[string]any{"field": "brand", "filter": "{}"}, nil)
			}},
		{"GetCollectionIndexes", http.MethodGet, "/v2/database/collections/products/indexes", url.Values{"databaseIntegrationId": {"int_1"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetCollectionIndexes(ctx, "products", map[string]any{"databaseIntegrationId": "int_1"}, nil)
			}},
		{"DeleteManyRecords", http.MethodDelete, "/v2/database/collections/products/many", url.Values{"filter": {"{}"}, "allRecords": {"true"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteManyRecords(ctx, "products", map[string]any{"filter": "{}", "allRecords": true}, nil)
			}},
		{"InsertManyRecords", http.MethodPost, "/v2/database/collections/products/many", nil, map[string]any{"documents": "[{\"title\":\"A\"},{\"title\":\"B\"}]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.InsertManyRecords(ctx, "products", map[string]any{"documents": "[{\"title\":\"A\"},{\"title\":\"B\"}]"}, nil)
			}},
		{"UpdateManyRecords", http.MethodPut, "/v2/database/collections/products/many", nil, map[string]any{"filter": "{}", "allRecords": true, "update": "{\"lines.$[line].qty\":3}", "arrayFilters": "[{\"line.sku\":\"A-1\"}]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateManyRecords(ctx, "products", map[string]any{"filter": "{}", "allRecords": true, "update": "{\"lines.$[line].qty\":3}", "arrayFilters": "[{\"line.sku\":\"A-1\"}]"}, nil)
			}},
		{"DeleteRecord", http.MethodDelete, "/v2/database/collections/products/rec_1", url.Values{"databaseIntegrationId": {"int_1"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteRecord(ctx, "products", "rec_1", map[string]any{"databaseIntegrationId": "int_1"}, nil)
			}},
		{"FindOneRecord", http.MethodGet, "/v2/database/collections/products/rec_1", url.Values{"databaseIntegrationId": {"int_1"}, "expandReferences": {"true"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindOneRecord(ctx, "products", "rec_1", map[string]any{"databaseIntegrationId": "int_1", "expandReferences": true}, nil)
			}},
		{"UpdateOneRecord", http.MethodPut, "/v2/database/collections/products/rec_1", nil, map[string]any{"update": "{\"lines.$[line].qty\":3}", "arrayFilters": "[{\"line.sku\":\"A-1\"}]"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateOneRecord(ctx, "products", "rec_1", map[string]any{"update": "{\"lines.$[line].qty\":3}", "arrayFilters": "[{\"line.sku\":\"A-1\"}]"}, nil)
			}},
		{"ReplaceRecord", http.MethodPut, "/v2/database/collections/products/rec_1/replace", nil, map[string]any{"replacement": "{\"title\":\"Boot\"}"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ReplaceRecord(ctx, "products", "rec_1", map[string]any{"replacement": "{\"title\":\"Boot\"}"}, nil)
			}},
		{"ChangeRecordResponsibility", http.MethodPut, "/v2/database/collections/products/rec_1/responsibility", nil, map[string]any{"newResponsibleUserId": "usr_2"},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ChangeRecordResponsibility(ctx, "products", "rec_1", map[string]any{"newResponsibleUserId": "usr_2"}, nil)
			}},
		{"ApplyDatabaseSchemaBundle", http.MethodPost, "/v2/database/schemas/apply-bundle", nil, map[string]any{"bundleJson": "{\"schemas\":[]}", "tier": "Free", "publish": true},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ApplyDatabaseSchemaBundle(ctx, map[string]any{"bundleJson": "{\"schemas\":[]}", "tier": "Free", "publish": true}, nil)
			}},
		{"UpdateDatabaseSchemaEmbed", http.MethodPut, "/v2/database/schemas/id_1/embed", nil, map[string]any{"embed": map[string]any{"enabled": true, "fields": []any{"title"}, "perUser": false}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateDatabaseSchemaEmbed(ctx, "id_1", map[string]any{"embed": map[string]any{"enabled": true, "fields": []any{"title"}, "perUser": false}}, nil)
			}},
		{"GetDatabaseSchemaIndexStatus", http.MethodGet, "/v2/database/schemas/id_1/index-status", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchemaIndexStatus(ctx, "id_1", nil, nil)
			}},
		{"GetDatabaseSchemaListSettings", http.MethodGet, "/v2/database/schemas/id_1/list-settings", nil, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseSchemaListSettings(ctx, "id_1", nil, nil)
			}},
		{"UpdateDatabaseSchemaListSettings", http.MethodPut, "/v2/database/schemas/id_1/list-settings", nil, map[string]any{"settings": map[string]any{"columns": []any{map[string]any{"field": "title"}}, "defaultSort": map[string]any{"field": "title"}}},
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateDatabaseSchemaListSettings(ctx, "id_1", map[string]any{"settings": map[string]any{"columns": []any{map[string]any{"field": "title"}}, "defaultSort": map[string]any{"field": "title"}}}, nil)
			}},
		{"GetDatabaseTaxonomyTree", http.MethodGet, "/v2/database/taxonomies/tree", url.Values{"includeTerms": {"true"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseTaxonomyTree(ctx, map[string]any{"includeTerms": true}, nil)
			}},
		{"GetDatabaseMergedTermTree", http.MethodGet, "/v2/database/taxonomies/categories/merged-tree", url.Values{"databaseIntegrationId": {"int_1"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseMergedTermTree(ctx, "categories", map[string]any{"databaseIntegrationId": "int_1"}, nil)
			}},
		{"GetDatabaseTaxonomyTermTree", http.MethodGet, "/v2/database/taxonomies/categories/terms/tree", url.Values{"rootTermId": {"t_1"}, "depth": {"2"}}, nil,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseTaxonomyTermTree(ctx, "categories", map[string]any{"rootTermId": "t_1", "depth": 2}, nil)
			}},
	}
}

func newHubDatabaseTestModule(baseURL string) *DatabaseModule {
	return &DatabaseModule{t: transport.New(&transport.Config{
		ProjectID:  "proj_1",
		AccountID:  "acct_1",
		APIKey:     "key_1",
		BaseURLAPI: "http://api.invalid",
		BaseURLHub: baseURL,
		APIVersion: "v2",
		HubVersion: "v2",
		Timeout:    5 * time.Second,
		Env:        "PROD",
	}, nil)}
}

func TestHubDatabaseEndpointsHitTheExpectedRoute(t *testing.T) {
	for _, c := range hubDatabaseCases() {
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

			if err := c.call(context.Background(), newHubDatabaseTestModule(srv.URL)); err != nil {
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
			assertFields(t, got, c.body)
		})
	}
}

// The HUB database surface is 67 routes (imports excluded). If the gateway
// grows one and a method is added, this count changes, so a new endpoint
// cannot arrive untested.
func TestHubDatabaseSurfaceSize(t *testing.T) {
	const want = 68
	if got := len(hubDatabaseCases()); got != want {
		t.Errorf("database endpoint count: got %d want %d", got, want)
	}
}

// A typed embed setting reaches the wire under "embed" with the gateway's
// field names, and the schema id goes in the path, not the body.
func TestUpdateDatabaseSchemaEmbedSendsTheTypedSettings(t *testing.T) {
	var gotPath string
	var got map[string]any
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotPath = r.URL.Path
		b, _ := io.ReadAll(r.Body)
		_ = json.Unmarshal(b, &got)
		w.WriteHeader(http.StatusNoContent)
	}))
	defer srv.Close()

	m := newHubDatabaseTestModule(srv.URL)
	req := map[string]any{"embed": dtos.SchemaEmbedSettingsDto{
		Enabled:                true,
		Fields:                 []string{"title", "description"},
		EmbeddingIntegrationId: "emb_1",
	}}
	if err := m.UpdateDatabaseSchemaEmbed(context.Background(), "sch_1", req, nil); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if gotPath != "/v2/database/schemas/sch_1/embed" {
		t.Errorf("path: got %q", gotPath)
	}
	assertFields(t, got, map[string]any{"embed": map[string]any{
		"enabled":                true,
		"fields":                 []any{"title", "description"},
		"embeddingIntegrationId": "emb_1",
	}})
	if _, ok := got["id"]; ok {
		t.Errorf("body: schema id must not be sent in the body, got %v", got["id"])
	}
}

// The taxonomy list rows decode Description, Dependencies, ParentName and
// DependencyRefs: one {id, name} pair per dependency, in the order of
// Dependencies. A dependency that no longer exists keeps its place with a null
// name, which decodes to an empty Name.
func TestGetDatabaseTaxonomiesDecodesTheDependencyRefs(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"result":[{"viewId":"tax_1","taxonomyName":"categories","parentId":"tax_0","parentName":"root","description":"Product categories","dependencies":["tax_2","tax_gone"],"dependencyRefs":[{"id":"tax_2","name":"brands"},{"id":"tax_gone","name":null}]}],"totalCount":1}`))
	}))
	defer srv.Close()

	var out struct {
		Result []dtos.TaxonomyListProjection `json:"result"`
	}
	if err := newHubDatabaseTestModule(srv.URL).GetDatabaseTaxonomies(context.Background(), nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(out.Result) != 1 {
		t.Fatalf("result count: got %d want 1", len(out.Result))
	}
	got, _ := json.Marshal(out.Result[0])
	want := `{"viewId":"tax_1","taxonomyName":"categories","parentId":"tax_0","description":"Product categories","dependencies":["tax_2","tax_gone"],"parentName":"root","dependencyRefs":[{"id":"tax_2","name":"brands"},{"id":"tax_gone"}]}`
	if string(got) != want {
		t.Errorf("taxonomy:\n got %s\nwant %s", got, want)
	}
}

// A schema trigger read carries its environment, and SchemaId is the owning
// schema (sch_...), not the trigger's own id. The list rows carry Env too.
func TestSchemaTriggerReadsDecodeEnvAndOwningSchema(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.URL.Path == "/v2/database/schemas/triggers" {
			_, _ = w.Write([]byte(`{"result":[{"id":"trg_1","name":"On insert","type":"Insert","env":"DEV"}],"totalCount":1}`))
			return
		}
		_, _ = w.Write([]byte(`{"result":{"id":"trg_1","name":"On insert","schemaId":"sch_1","when":"Insert","env":"DEV"}}`))
	}))
	defer srv.Close()
	m := newHubDatabaseTestModule(srv.URL)

	var one struct {
		Result dtos.SchemaTriggerDto `json:"result"`
	}
	if err := m.GetSchemaTrigger(context.Background(), "trg_1", nil, &one); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if one.Result.SchemaId != "sch_1" || one.Result.Env != "DEV" {
		t.Errorf("trigger: got schemaId %q env %q, want sch_1 / DEV", one.Result.SchemaId, one.Result.Env)
	}

	var list struct {
		Result []dtos.SchemaTriggerProjectionList `json:"result"`
	}
	if err := m.GetSchemaTriggers(context.Background(), nil, &list); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(list.Result) != 1 || list.Result[0].Env != "DEV" {
		t.Errorf("trigger list: got %+v, want one row with env DEV", list.Result)
	}
}

// The schema trigger calls act on the copy in the request environment, so the
// client's environment must reach the gateway as the norbix-env header.
func TestSchemaTriggerCallsSendTheEnvironment(t *testing.T) {
	calls := map[string]func(ctx context.Context, m *DatabaseModule) error{
		"GetSchemaTriggers": func(ctx context.Context, m *DatabaseModule) error { return m.GetSchemaTriggers(ctx, nil, nil) },
		"EnableSchemaTrigger": func(ctx context.Context, m *DatabaseModule) error {
			return m.EnableSchemaTrigger(ctx, "trg_1", nil, nil)
		},
		"DisableSchemaTrigger": func(ctx context.Context, m *DatabaseModule) error {
			return m.DisableSchemaTrigger(ctx, "trg_1", nil, nil)
		},
		"DeleteSchemaTrigger": func(ctx context.Context, m *DatabaseModule) error {
			return m.DeleteSchemaTrigger(ctx, "trg_1", nil, nil)
		},
	}
	for name, call := range calls {
		t.Run(name, func(t *testing.T) {
			var gotEnv string
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				gotEnv = r.Header.Get("norbix-env")
				w.WriteHeader(http.StatusNoContent)
			}))
			defer srv.Close()
			m := &DatabaseModule{t: transport.New(&transport.Config{
				ProjectID: "proj_1", AccountID: "acct_1", APIKey: "key_1",
				BaseURLAPI: "http://api.invalid", BaseURLHub: srv.URL,
				APIVersion: "v2", HubVersion: "v2", Timeout: 5 * time.Second, Env: "DEV",
			}, nil)}
			if err := call(context.Background(), m); err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if gotEnv != "DEV" {
				t.Errorf("norbix-env: got %q want DEV", gotEnv)
			}
		})
	}
}

// A saved aggregate names the collections its pipeline joins.
func TestGetDatabaseAggregateDecodesJoinedCollections(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(`{"result":{"viewId":"agg_1","displayName":"Orders with users","schemaViewId":"sch_1","pipeline":"[]","joinedCollections":["users","products"]}}`))
	}))
	defer srv.Close()

	var out struct {
		Result dtos.MongoDbAggregateDto `json:"result"`
	}
	if err := newHubDatabaseTestModule(srv.URL).GetDatabaseAggregate(context.Background(), "agg_1", nil, &out); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	got, _ := json.Marshal(out.Result.JoinedCollections)
	if string(got) != `["users","products"]` {
		t.Errorf("joinedCollections: got %s", got)
	}
}

// The Database refusals added on refactoringV2 reach the caller with the
// gateway's own code and context, so callers can switch on them.
func TestDatabaseRefusalsKeepTheGatewayCode(t *testing.T) {
	cases := []struct {
		name    string
		status  int
		code    string
		context string // JSON of the error's context, "" for none
		call    func(ctx context.Context, m *DatabaseModule) error
	}{
		{"empty filter without allRecords", http.StatusBadRequest, "CM-ERRORS-DATABASE-037", `{"Operation":"delete"}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteManyRecords(ctx, "products", map[string]any{"filter": "{}"}, nil)
			}},
		{"update with a $ operator", http.StatusBadRequest, "CM-ERRORS-DATABASE-035", `{"Operator":"$inc"}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.UpdateOneRecord(ctx, "products", "rec_1", map[string]any{"update": `{"$inc":{"stock":1}}`}, nil)
			}},
		{"insert many with a broken item", http.StatusBadRequest, "CM-ERRORS-DATABASE-036", `{"Index":1}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.InsertManyRecords(ctx, "products", map[string]any{"documents": `[{"title":"A"},3]`}, nil)
			}},
		{"new owner not a project user", http.StatusBadRequest, "CM-ERRORS-MEMBERSHIP-USERS-012", "",
			func(ctx context.Context, m *DatabaseModule) error {
				return m.ChangeRecordResponsibility(ctx, "products", "rec_1", map[string]any{"newResponsibleUserId": "usr_x"}, nil)
			}},
		{"schema used by a saved aggregate", http.StatusBadRequest, "CM-ERRORS-SCHEMA-018", `{"BlockerAggregateIds":["agg_1"],"BlockerAggregateNames":["Orders with users"],"SchemaId":"sch_1"}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.DeleteDatabaseSchema(ctx, "sch_1", nil, nil)
			}},
		{"rename to a used name", http.StatusBadRequest, "CM-ERRORS-SCHEMA-002", `{"SchemaName":"orders"}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.RenameDatabaseSchema(ctx, "sch_1", map[string]any{"title": "orders"}, nil)
			}},
		{"unknown taxonomy name", http.StatusNotFound, "CM-ERRORS-TAXONOMIES-010", "",
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseMergedTermTree(ctx, "nope", nil, nil)
			}},
		{"term tree too large", http.StatusBadRequest, "CM-ERRORS-TAXONOMIES-011", `{"MaxTerms":5000}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.GetDatabaseTaxonomyTree(ctx, map[string]any{"includeTerms": true}, nil)
			}},
		{"expand references: a linked source is not readable", http.StatusForbidden, "CM-ERRORS-DATABASE-056", `{"Fields":["owner"],"MissingPermissions":["membership:read"],"Source":"users","SourceKind":"user"}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.FindRecords(ctx, "products", map[string]any{"expandReferences": true}, nil)
			}},
		{"reference to a missing record", http.StatusBadRequest, "CM-ERRORS-DATABASE-053", `{"FieldName":"category","Keyword":"reference","MissingId":"rec_x"}`,
			func(ctx context.Context, m *DatabaseModule) error {
				return m.InsertRecord(ctx, "products", map[string]any{"document": `{"category":"rec_x"}`}, nil)
			}},
		{"trigger copy missing in env", http.StatusNotFound, "CM-ERRORS-TRIGGERS-002", "",
			func(ctx context.Context, m *DatabaseModule) error {
				return m.EnableSchemaTrigger(ctx, "trg_1", nil, nil)
			}},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			entry := map[string]any{"errorCode": c.code, "message": "refused"}
			if c.context != "" {
				var ctxMap map[string]any
				_ = json.Unmarshal([]byte(c.context), &ctxMap)
				entry["context"] = ctxMap
			}
			body, _ := json.Marshal(map[string]any{"responseStatus": map[string]any{"isSuccess": false, "errors": []any{entry}}})
			srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(c.status)
				_, _ = w.Write(body)
			}))
			defer srv.Close()

			err := c.call(context.Background(), newHubDatabaseTestModule(srv.URL))
			var nbErr *nberrors.Error
			if !errors.As(err, &nbErr) {
				t.Fatalf("error: got %T %v, want *errors.Error", err, err)
			}
			if nbErr.Code != c.code {
				t.Errorf("code: got %q want %q", nbErr.Code, c.code)
			}
			if c.context != "" {
				if len(nbErr.Errors) != 1 {
					t.Fatalf("errors: got %d want 1", len(nbErr.Errors))
				}
				got, _ := json.Marshal(nbErr.Errors[0].Context)
				if string(got) != c.context {
					t.Errorf("context: got %s want %s", got, c.context)
				}
			}
		})
	}
}
