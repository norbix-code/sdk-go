package hub

// ptr points at v, for the optional (*bool, *float64) fields of the generated
// request types. Apps use norbix.Bool / norbix.Float64; this package cannot
// import norbix (norbix imports hub).
func ptr[T any](v T) *T { return &v }
