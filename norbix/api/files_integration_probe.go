// Hand-written, not generated (10b-files slice API-TEST, issue #39).
//
// files.go carries the generator's DO NOT EDIT header, so this method lives in
// its own file of the same package. If a future regeneration of files.go adds
// TestFilesIntegration itself, the build fails with "method redeclared" and
// this file should simply be deleted.

package api

import (
	"context"

	"github.com/norbix-code/sdk-go/norbix/internal/transport"
)

// TestFilesIntegration performs POST /{version}/files/{filesIntegrationId}/test (scope: project).
//
// It runs a live probe against the files integration: it uploads a small
// file, reads it back, lists the folder and deletes the file again. The
// answer has one item per step, in order UploadFile, GetFile, GetAllFiles,
// DeleteFile, each with an Operation, a Result ("OK", "FAILED", or
// "NOT_TESTED" once an earlier step failed) and, when the step failed, its
// Errors. Decode it into a
// *dtos.TestFilesIntegrationResponse:
//
//	var res dtos.TestFilesIntegrationResponse
//	err := client.API.Files.TestFilesIntegration(ctx, integrationID, nil, &res)
//
// Because the probe writes to the storage, the gateway asks for the
// files:create permission. This is the public API endpoint; the dashboard's
// own endpoint is hub.FilesModule.TestFilesIntegration
// (POST /{version}/files/integrations/test).
func (m *FilesModule) TestFilesIntegration(ctx context.Context, filesIntegrationId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"filesIntegrationId": filesIntegrationId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/files/{filesIntegrationId}/test",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}
