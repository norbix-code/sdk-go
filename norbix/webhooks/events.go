package webhooks

// EventNames is the closed catalog of event names a destination may subscribe to.
// Source: gateway Domain trigger event name value objects.
var EventNames = []string{
	"database.record.inserted",
	"database.record.updated",
	"database.record.deleted",
	"database.record.replaced",
	"database.record.responsibilityChanged",
	"database.records.inserted",
	"database.records.updated",
	"database.records.deleted",
	"membership.user.registered",
	"membership.user.invited",
	"membership.user.verified",
	"membership.user.updated",
	"membership.user.deleted",
	"membership.user.blocked",
	"membership.user.reactivated",
	"files.file.uploaded",
	"files.file.deleted",
	"payments.webhook.received",
}

// Named event constants — use these instead of raw strings.
const (
	EventDatabaseRecordInserted              = "database.record.inserted"
	EventDatabaseRecordUpdated               = "database.record.updated"
	EventDatabaseRecordDeleted               = "database.record.deleted"
	EventDatabaseRecordReplaced              = "database.record.replaced"
	EventDatabaseRecordResponsibilityChanged = "database.record.responsibilityChanged"
	EventDatabaseRecordsInserted             = "database.records.inserted"
	EventDatabaseRecordsUpdated              = "database.records.updated"
	EventDatabaseRecordsDeleted              = "database.records.deleted"
	EventMembershipUserRegistered            = "membership.user.registered"
	EventMembershipUserInvited               = "membership.user.invited"
	EventMembershipUserVerified              = "membership.user.verified"
	EventMembershipUserUpdated               = "membership.user.updated"
	EventMembershipUserDeleted               = "membership.user.deleted"
	EventMembershipUserBlocked               = "membership.user.blocked"
	EventMembershipUserReactivated           = "membership.user.reactivated"
	EventFilesFileUploaded                   = "files.file.uploaded"
	EventFilesFileDeleted                    = "files.file.deleted"
	// EventPaymentsWebhookReceived: a payment provider (Stripe, Apple In-App,
	// Google In-App, …) called the project's payments webhook. The payload is the
	// normalized payment fact (decode into dtos.PaymentFactDto); the provider's
	// own event name is its EventName.
	EventPaymentsWebhookReceived = "payments.webhook.received"
)
