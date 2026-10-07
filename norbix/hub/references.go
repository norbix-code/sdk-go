package hub

// ExpandedReference is what a reference field (user, role, taxonomy term,
// record, file) holds in a record read with "expandReferences": true —
// a multiple field holds a list of them. The gateway builds it on the fly
// (no generated DTO), so it lives here next to the module.
//
// Display is the target's property the schema's displayField names, read
// in the request environment: a string most of the time, but it keeps the
// property's own JSON type. It is nil when the target no longer exists;
// Id is always set.
type ExpandedReference struct {
	Id      string `json:"id"`
	Display any    `json:"display"`
}
