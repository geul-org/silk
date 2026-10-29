//ff:type feature=scan type=model
//ff:what Matcher — tests one record's bytes (TAG WORD first); must be safe for concurrent use
package scan

// Matcher is the predicate of a scan. rec is exactly one record and is only
// valid during the call.
type Matcher interface {
	Match(rec []byte) bool
}
