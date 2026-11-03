// Hand-written, not generated (gateway campaign audit/messaging, owner decision
// 2026-11-03).
//
// Written in the shape of the other modules (payments.go, hub/notifications.go),
// whose generator was lost. The request and response types are the generated
// dtos (CreateEmailCampaignRequest, GetOwnPushMessagesResponse, …); every path
// is copied verbatim from sdks/typegen/coverage/endpoints.api.json.

package api

import (
	"context"

	"github.com/norbix-code/sdk-go/v2/norbix/internal/transport"
)

// NotificationsModule groups the notification endpoints on the API: Email, SMS
// and Push campaigns (create, read, batches, statistics, stop, retry the
// failed ones, preview) and the signed-in user's own push devices and
// messages. The Hub has the same campaign paths for the admin surface
// (c.Hub.Notifications); this module is the data-API twin a project's own
// backend or app calls with an API key or the user's bearer token.
type NotificationsModule struct{ t *transport.Transport }

// RegisterOwnPushDevice performs POST /{version}/notifications/push/devices (scope: project).
// Registers or refreshes the signed-in user's push device. Body: provider (fcm
// or apns — an Expo push token is refused), token, optional platform, osName,
// osVersion, appVersion, deviceName, modelName, brand, manufacturer,
// installationId, locale, timeZone. The same token again answers the same id; a
// token another user held moves to the caller. Answers {id: "pnd_…"}.
// Request dtos.RegisterOwnPushDeviceRequest; answers dtos.IdResponse.
func (m *NotificationsModule) RegisterOwnPushDevice(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/devices",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetOwnPushDevices performs GET /{version}/notifications/push/devices (scope: project).
// Lists the caller's own push devices, last refreshed first. Answers {items:
// [...]} (dtos.GetOwnPushDevicesResponse).
// Request dtos.GetOwnPushDevicesRequest; answers dtos.GetOwnPushDevicesResponse.
func (m *NotificationsModule) GetOwnPushDevices(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/devices",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// DeleteOwnPushDevice performs DELETE /{version}/notifications/push/devices/{Id} (scope: project).
// Removes one of the caller's own devices; another user's device is not found.
// Request dtos.DeleteOwnPushDeviceRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) DeleteOwnPushDevice(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/devices/{Id}",
		Method:     "DELETE",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// DeleteOwnPushDeviceByToken performs DELETE /{version}/notifications/push/devices (scope: project).
// Sign-out from a device: removes the caller's device registered with query
// token. Safe to repeat.
// Request dtos.DeleteOwnPushDeviceByTokenRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) DeleteOwnPushDeviceByToken(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/devices",
		Method:     "DELETE",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetOwnPushMessages performs GET /{version}/notifications/push/messages (scope: project).
// Lists the caller's own push messages, newest first, paged (pageNumber,
// pageSize); unreadOnly=true keeps the unread ones. Answers {list: {items,
// totalCount, …}} (dtos.GetOwnPushMessagesResponse).
// Request dtos.GetOwnPushMessagesRequest; answers dtos.GetOwnPushMessagesResponse.
func (m *NotificationsModule) GetOwnPushMessages(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/messages",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// MarkOwnPushMessageRead performs POST /{version}/notifications/push/messages/{Id}/read (scope: project).
// Marks one of the caller's own push messages read. Safe to repeat; another
// user's message is not found.
// Request dtos.MarkOwnPushMessageReadRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) MarkOwnPushMessageRead(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/messages/{Id}/read",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// MarkAllOwnPushMessagesRead performs POST /{version}/notifications/push/messages/read (scope: project).
// Marks all of the caller's unread push messages read.
// Request dtos.MarkAllOwnPushMessagesReadRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) MarkAllOwnPushMessagesRead(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/messages/read",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// CreateEmailCampaign performs POST /{version}/notifications/email/campaigns (scope: project).
// Body: campaign (dtos.EmailCampaignRequest), optional databaseIntegrationId and
// sendNow. Answers {id}.
// Request dtos.CreateEmailCampaignRequest; answers dtos.IdResponse.
func (m *NotificationsModule) CreateEmailCampaign(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// DeleteEmailCampaign performs DELETE /{version}/notifications/email/campaigns/{Id} (scope: project).
// Request dtos.DeleteEmailCampaignRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) DeleteEmailCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{Id}",
		Method:     "DELETE",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEmailCampaign performs GET /{version}/notifications/email/campaigns/{id} (scope: project).
// Request dtos.GetEmailCampaign; answers dtos.GetEmailCampaignResponse.
func (m *NotificationsModule) GetEmailCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{id}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEmailCampaigns performs GET /{version}/notifications/email/campaigns (scope: project).
// Optional query: campaignId, emailAddress, templateId, from, to,
// databaseIntegrationId and the page (pageNumber, pageSize).
// Request dtos.GetEmailCampaigns; answers dtos.GetEmailCampaignsResponse.
func (m *NotificationsModule) GetEmailCampaigns(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEmailCampaignBatches performs GET /{version}/notifications/email/campaigns/{id}/batches (scope: project).
// Optional query: batchId, emailAddress and the page.
// Request dtos.GetEmailCampaignBatches; answers
// dtos.GetEmailCampaignBatchesResponse.
func (m *NotificationsModule) GetEmailCampaignBatches(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{id}/batches",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEmailCampaignBatchNotification performs GET /{version}/notifications/email/campaigns/{id}/batches/{batchId}/{notificationId} (scope: project).
// Answers one campaign notification row with its delivery state, attempts and
// delivery events.
// Request dtos.GetEmailCampaignBatchNotification; answers
// dtos.GetEmailCampaignBatchNotificationResponse.
func (m *NotificationsModule) GetEmailCampaignBatchNotification(ctx context.Context, id string, batchId string, notificationId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id":             id,
		"batchId":        batchId,
		"notificationId": notificationId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{id}/batches/{batchId}/{notificationId}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEmailCampaignBatchNotifications performs GET /{version}/notifications/email/campaigns/{id}/batches/{batchId} (scope: project).
// Request dtos.GetEmailCampaignBatchNotifications; answers
// dtos.GetEmailCampaignBatchNotificationsResponse.
func (m *NotificationsModule) GetEmailCampaignBatchNotifications(ctx context.Context, id string, batchId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id":      id,
		"batchId": batchId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{id}/batches/{batchId}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEmailCampaignStatistics performs GET /{version}/notifications/email/campaigns/{id}/stats (scope: project).
// Request dtos.GetEmailCampaignStatistics; answers
// dtos.GetEmailCampaignStatisticsResponse.
func (m *NotificationsModule) GetEmailCampaignStatistics(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{id}/stats",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// PreviewEmailNotification performs GET /{version}/notifications/email/preview (scope: optional — a signed link in "hash" opens it without sign-in).
// The hash alone is the key: a client with no API key and no bearer token still
// sends the request. A signed-in member may pass projectId + notificationId
// instead.
// Request dtos.PreviewEmailNotification; answers
// dtos.PreviewEmailNotificationResponse.
func (m *NotificationsModule) PreviewEmailNotification(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/preview",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeOptional,
	}, out)
}

// RetryFailedEmailNotifications performs POST /{version}/notifications/email/campaigns/{Id}/retry-failed (scope: project).
// Re-sends the campaign's retryable failures. Body may carry notificationIds,
// categories, includeAmbiguous and an idempotencyKey; the answer
// (dtos.RetryFailedEmailNotificationsResponse) says how many were selected and
// skipped.
// Request dtos.RetryFailedEmailNotificationsRequest; answers
// dtos.RetryFailedEmailNotificationsResponse.
func (m *NotificationsModule) RetryFailedEmailNotifications(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{Id}/retry-failed",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// StopEmailCampaign performs POST /{version}/notifications/email/campaigns/{Id}/stop (scope: project).
// Request dtos.StopEmailCampaignRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) StopEmailCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/email/campaigns/{Id}/stop",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetEmailCampaignMessages performs GET /{version}/notifications/emails/campaigns/{campaignId}/messages (scope: project).
// Note the "emails" segment: this is the only route under
// /notifications/emails/. Optional query: campaignBatchId and the page.
// Request dtos.GetEmailCampaignMessagesRequest; answers
// dtos.GetEmailCampaignMessagesResponse.
func (m *NotificationsModule) GetEmailCampaignMessages(ctx context.Context, campaignId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"campaignId": campaignId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/emails/campaigns/{campaignId}/messages",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// CreateSmsCampaign performs POST /{version}/notifications/sms/campaigns (scope: project).
// Body: templateId, deliveryType and the matching audience (allUsers,
// specifiedUsers, accountUsers, collection, phoneNumbers), optional
// integrationId, language, initiatorId, databaseIntegrationId and sendNow.
// Answers {id}.
// Request dtos.CreateSmsCampaignRequest; answers dtos.IdResponse.
func (m *NotificationsModule) CreateSmsCampaign(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// DeleteSmsCampaign performs DELETE /{version}/notifications/sms/campaigns/{id} (scope: project).
// Request dtos.DeleteSmsCampaign; answers nothing (EmptyResponse).
func (m *NotificationsModule) DeleteSmsCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{id}",
		Method:     "DELETE",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetSmsCampaign performs GET /{version}/notifications/sms/campaigns/{id} (scope: project).
// Request dtos.GetSmsCampaign; answers dtos.GetSmsCampaignResponse.
func (m *NotificationsModule) GetSmsCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{id}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetSmsCampaigns performs GET /{version}/notifications/sms/campaigns (scope: project).
// Optional query: campaignId, templateId, from, to, databaseIntegrationId and
// the page.
// Request dtos.GetSmsCampaigns; answers dtos.GetSmsCampaignsResponse.
func (m *NotificationsModule) GetSmsCampaigns(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetSmsCampaignBatches performs GET /{version}/notifications/sms/campaigns/{id}/batches (scope: project).
// Request dtos.GetSmsCampaignBatches; answers dtos.GetSmsCampaignBatchesResponse.
func (m *NotificationsModule) GetSmsCampaignBatches(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{id}/batches",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetSmsCampaignBatchNotification performs GET /{version}/notifications/sms/campaigns/{id}/batches/{batchId}/{notificationId} (scope: project).
// Request dtos.GetSmsCampaignBatchNotification; answers
// dtos.GetSmsCampaignBatchNotificationResponse.
func (m *NotificationsModule) GetSmsCampaignBatchNotification(ctx context.Context, id string, batchId string, notificationId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id":             id,
		"batchId":        batchId,
		"notificationId": notificationId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{id}/batches/{batchId}/{notificationId}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetSmsCampaignBatchNotifications performs GET /{version}/notifications/sms/campaigns/{id}/batches/{batchId} (scope: project).
// Request dtos.GetSmsCampaignBatchNotifications; answers
// dtos.GetSmsCampaignBatchNotificationsResponse.
func (m *NotificationsModule) GetSmsCampaignBatchNotifications(ctx context.Context, id string, batchId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id":      id,
		"batchId": batchId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{id}/batches/{batchId}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetSmsCampaignStatistics performs GET /{version}/notifications/sms/campaigns/{id}/stats (scope: project).
// Request dtos.GetSmsCampaignStatistics; answers
// dtos.GetSmsCampaignStatisticsResponse.
func (m *NotificationsModule) GetSmsCampaignStatistics(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{id}/stats",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// PreviewSmsNotification performs GET /{version}/notifications/sms/preview (scope: optional — a signed link in "hash" opens it without sign-in).
// The hash alone is the key: a client with no API key and no bearer token still
// sends the request. A signed-in member may pass projectId + notificationId
// instead.
// Request dtos.PreviewSmsNotification; answers dtos.PreviewSmsNotificationResponse.
func (m *NotificationsModule) PreviewSmsNotification(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/preview",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeOptional,
	}, out)
}

// RetryFailedSmsNotifications performs POST /{version}/notifications/sms/campaigns/{Id}/retry-failed (scope: project).
// Same body and answer shape as RetryFailedEmailNotifications
// (dtos.RetryFailedSmsNotificationsResponse).
// Request dtos.RetryFailedSmsNotificationsRequest; answers
// dtos.RetryFailedSmsNotificationsResponse.
func (m *NotificationsModule) RetryFailedSmsNotifications(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{Id}/retry-failed",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// StopSmsCampaign performs POST /{version}/notifications/sms/campaigns/{Id}/stop (scope: project).
// Request dtos.StopSmsCampaignRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) StopSmsCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{Id}/stop",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetSmsCampaignMessages performs GET /{version}/notifications/sms/campaigns/{campaignId}/messages (scope: project).
// Optional query: campaignBatchId and the page.
// Request dtos.GetSmsCampaignMessagesRequest; answers
// dtos.GetSmsCampaignMessagesResponse.
func (m *NotificationsModule) GetSmsCampaignMessages(ctx context.Context, campaignId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"campaignId": campaignId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/sms/campaigns/{campaignId}/messages",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaignAudienceCount performs POST /{version}/notifications/push/campaigns/audience-count (scope: project).
// Counts who a campaign would reach before it is created. Body: campaign
// (dtos.PushCampaignRequest), optional databaseIntegrationId. Answers
// dtos.GetPushCampaignAudienceCountResponse.
// Request dtos.GetPushCampaignAudienceCountRequest; answers
// dtos.GetPushCampaignAudienceCountResponse.
func (m *NotificationsModule) GetPushCampaignAudienceCount(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/audience-count",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// CreatePushCampaign performs POST /{version}/notifications/push/campaigns (scope: project).
// Body: campaign (dtos.PushCampaignRequest), optional databaseIntegrationId and
// sendNow. Answers {id}.
// Request dtos.CreatePushCampaignRequest; answers dtos.IdResponse.
func (m *NotificationsModule) CreatePushCampaign(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns",
		Method:     "POST",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// DeletePushCampaign performs DELETE /{version}/notifications/push/campaigns/{Id} (scope: project).
// Request dtos.DeletePushCampaignRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) DeletePushCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{Id}",
		Method:     "DELETE",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaign performs GET /{version}/notifications/push/campaigns/{id} (scope: project).
// Request dtos.GetPushCampaign; answers dtos.GetPushCampaignResponse.
func (m *NotificationsModule) GetPushCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{id}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaigns performs GET /{version}/notifications/push/campaigns (scope: project).
// Optional query: templateId, from, to, databaseIntegrationId and the page.
// Request dtos.GetPushCampaigns; answers dtos.GetPushCampaignsResponse.
func (m *NotificationsModule) GetPushCampaigns(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaignBatches performs GET /{version}/notifications/push/campaigns/{id}/batches (scope: project).
// Optional query: batchId and the page.
// Request dtos.GetPushCampaignBatches; answers dtos.GetPushCampaignBatchesResponse.
func (m *NotificationsModule) GetPushCampaignBatches(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{id}/batches",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaignBatchNotification performs GET /{version}/notifications/push/campaigns/{id}/batches/{batchId}/{notificationId} (scope: project).
// Request dtos.GetPushCampaignBatchNotification; answers
// dtos.GetPushCampaignBatchNotificationResponse.
func (m *NotificationsModule) GetPushCampaignBatchNotification(ctx context.Context, id string, batchId string, notificationId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id":             id,
		"batchId":        batchId,
		"notificationId": notificationId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{id}/batches/{batchId}/{notificationId}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaignBatchNotifications performs GET /{version}/notifications/push/campaigns/{id}/batches/{batchId} (scope: project).
// Request dtos.GetPushCampaignBatchNotifications; answers
// dtos.GetPushCampaignBatchNotificationsResponse.
func (m *NotificationsModule) GetPushCampaignBatchNotifications(ctx context.Context, id string, batchId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id":      id,
		"batchId": batchId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{id}/batches/{batchId}",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaignStatistics performs GET /{version}/notifications/push/campaigns/{id}/stats (scope: project).
// Request dtos.GetPushCampaignStatistics; answers
// dtos.GetPushCampaignStatisticsResponse.
func (m *NotificationsModule) GetPushCampaignStatistics(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{id}/stats",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// PreviewPushNotification performs GET /{version}/notifications/push/preview (scope: optional — a signed link in "hash" opens it without sign-in).
// The hash alone is the key: a client with no API key and no bearer token still
// sends the request. A signed-in member may pass projectId + notificationId
// instead.
// Request dtos.PreviewPushNotification; answers
// dtos.PreviewPushNotificationResponse.
func (m *NotificationsModule) PreviewPushNotification(ctx context.Context, req map[string]any, out any) error {
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/preview",
		Method:     "GET",
		PathParams: nil,
		Body:       req,
		Scope:      transport.ScopeOptional,
	}, out)
}

// RetryFailedPushNotifications performs POST /{version}/notifications/push/campaigns/{Id}/retry-failed (scope: project).
// Same body and answer shape as RetryFailedEmailNotifications
// (dtos.RetryFailedPushNotificationsResponse).
// Request dtos.RetryFailedPushNotificationsRequest; answers
// dtos.RetryFailedPushNotificationsResponse.
func (m *NotificationsModule) RetryFailedPushNotifications(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{Id}/retry-failed",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// StopPushCampaign performs POST /{version}/notifications/push/campaigns/{Id}/stop (scope: project).
// Request dtos.StopPushCampaignRequest; answers nothing (EmptyResponse).
func (m *NotificationsModule) StopPushCampaign(ctx context.Context, id string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"Id": id,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{Id}/stop",
		Method:     "POST",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}

// GetPushCampaignMessages performs GET /{version}/notifications/push/campaigns/{campaignId}/messages (scope: project).
// Optional query: campaignBatchId and the page.
// Request dtos.GetPushCampaignMessagesRequest; answers
// dtos.GetPushCampaignMessagesResponse.
func (m *NotificationsModule) GetPushCampaignMessages(ctx context.Context, campaignId string, req map[string]any, out any) error {
	pathParams := map[string]string{
		"campaignId": campaignId,
	}
	return m.t.Send(ctx, transport.Request{
		Target:     transport.TargetAPI,
		Path:       "/{version}/notifications/push/campaigns/{campaignId}/messages",
		Method:     "GET",
		PathParams: pathParams,
		Body:       req,
		Scope:      transport.ScopeProject,
	}, out)
}
