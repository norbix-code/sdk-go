package norbix

// Optional fields of the generated types.
//
// A field the gateway reads as "not sent = unchanged" is a pointer in
// norbix/api/dtos and norbix/hub/dtos: *bool for a true/false setting,
// *float64 for a number. nil leaves the field out of the request; a pointer
// sends its value, false and 0 included:
//
//	settings := hubdtos.SaveLogSettings{
//		SkipHttpBodyMeta: norbix.Bool(false), // sent: switches the setting off
//		// SkipCloudDashboardLogs is nil: not sent, the stored value stays
//	}

// Bool returns a pointer to v, for an optional true/false field such as
// SaveLogSettings.SkipHttpBodyMeta. Bool(false) sends false.
func Bool(v bool) *bool { return &v }

// Float64 returns a pointer to v, for an optional number field such as
// GetPaymentsWebhookLog.Limit. Float64(0) sends 0.
func Float64(v float64) *float64 { return &v }
