package webhooks

import (
	"encoding/json"
	"strings"
)

// Normalized is the result of turning a wire envelope into a payload + metadata.
type Normalized struct {
	// Payload is the inner value handed to a handler, still JSON-encoded so the
	// caller can decode it into a concrete type (see DecodePayload / DecodeMutation).
	Payload  json.RawMessage
	Metadata EventMetadata
}

func rawField(data map[string]json.RawMessage, key string) json.RawMessage {
	if v, ok := data[key]; ok {
		return v
	}
	return nil
}

func rawString(data map[string]json.RawMessage, key string) string {
	v, ok := data[key]
	if !ok {
		return ""
	}
	var s string
	if json.Unmarshal(v, &s) == nil {
		return s
	}
	return ""
}

// Normalize turns a raw wire envelope into {payload, metadata}, mirroring the
// JS receiver:
//
//   - Entity events  -> payload is the entity (user / document / file).
//   - Mutation events -> payload is {from, to}.
//   - Batch events   -> payload is the array.
//
// Wrapper ids (record id, schema, user id, ...) are lifted onto Metadata.
// Unknown events fall back to payload = data, metadata = {}.
func Normalize(env Envelope) Normalized {
	event := env.Event
	var d map[string]json.RawMessage
	_ = json.Unmarshal(env.Data, &d)
	if d == nil {
		d = map[string]json.RawMessage{}
	}

	switch {
	case strings.HasPrefix(event, "database."):
		meta := EventMetadata{}
		if name := rawString(d, "schemaName"); name != "" {
			meta.SchemaName = name
			if schemaRaw := rawField(d, "schema"); schemaRaw != nil {
				var schema map[string]json.RawMessage
				if json.Unmarshal(schemaRaw, &schema) == nil {
					meta.SchemaID = rawString(schema, "id")
				}
			}
		}
		if iid := rawString(d, "integrationId"); iid != "" {
			meta.IntegrationID = iid
		}
		if id := rawString(d, "id"); id != "" {
			meta.RecordID = id
		}
		if idsRaw := rawField(d, "ids"); idsRaw != nil {
			var ids []string
			if json.Unmarshal(idsRaw, &ids) == nil {
				meta.RecordIDs = ids
			}
		}
		switch event {
		case EventDatabaseRecordInserted, EventDatabaseRecordDeleted:
			return Normalized{Payload: rawField(d, "document"), Metadata: meta}
		case EventDatabaseRecordUpdated, EventDatabaseRecordReplaced:
			return Normalized{Payload: fromTo(d), Metadata: meta}
		case EventDatabaseRecordsInserted:
			docs := rawField(d, "documents")
			if docs == nil {
				docs = json.RawMessage("[]")
			}
			return Normalized{Payload: docs, Metadata: meta}
		default:
			return Normalized{Payload: env.Data, Metadata: meta}
		}

	case strings.HasPrefix(event, "membership."):
		meta := EventMetadata{}
		if id := rawString(d, "id"); id != "" {
			meta.UserID = id
		}
		switch event {
		case EventMembershipUserRegistered, EventMembershipUserVerified,
			EventMembershipUserBlocked, EventMembershipUserReactivated:
			return Normalized{Payload: rawField(d, "to"), Metadata: meta}
		case EventMembershipUserDeleted:
			return Normalized{Payload: rawField(d, "from"), Metadata: meta}
		case EventMembershipUserUpdated:
			return Normalized{Payload: fromTo(d), Metadata: meta}
		case EventMembershipUserInvited:
			return Normalized{Payload: pick(d, "email"), Metadata: meta}
		default:
			return Normalized{Payload: env.Data, Metadata: meta}
		}

	case strings.HasPrefix(event, "files."):
		meta := EventMetadata{}
		if iid := rawString(d, "integrationId"); iid != "" {
			meta.IntegrationID = iid
		}
		switch event {
		case EventFilesFileUploaded:
			return Normalized{Payload: rawField(d, "file"), Metadata: meta}
		case EventFilesFileDeleted:
			return Normalized{Payload: pick(d, "path"), Metadata: meta}
		default:
			return Normalized{Payload: env.Data, Metadata: meta}
		}
	}

	return Normalized{Payload: env.Data, Metadata: EventMetadata{}}
}

func fromTo(d map[string]json.RawMessage) json.RawMessage {
	out := map[string]json.RawMessage{"from": rawOrNull(d, "from"), "to": rawOrNull(d, "to")}
	b, _ := json.Marshal(out)
	return b
}

func pick(d map[string]json.RawMessage, key string) json.RawMessage {
	out := map[string]json.RawMessage{key: rawOrNull(d, key)}
	b, _ := json.Marshal(out)
	return b
}

func rawOrNull(d map[string]json.RawMessage, key string) json.RawMessage {
	if v, ok := d[key]; ok && v != nil {
		return v
	}
	return json.RawMessage("null")
}

// DecodePayload decodes a handler payload into T.
func DecodePayload[T any](payload json.RawMessage) (T, error) {
	var v T
	if len(payload) == 0 {
		return v, nil
	}
	err := json.Unmarshal(payload, &v)
	return v, err
}

// DecodeMutation decodes a {from, to} payload into Mutation[T].
func DecodeMutation[T any](payload json.RawMessage) (Mutation[T], error) {
	var m Mutation[T]
	if len(payload) == 0 {
		return m, nil
	}
	err := json.Unmarshal(payload, &m)
	return m, err
}
