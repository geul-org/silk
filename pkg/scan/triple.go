//ff:type feature=scan type=model
//ff:what Triple — a TRIPLE Matcher: the Predicate and Object records' exact bytes (nil matches any)
package scan

// Triple matches a TRIPLE (edge/triple.md: Edge Id, Predicate, Subject,
// Object, Rank) whose Predicate and Object records equal Pred and Obj byte
// for byte. GEUL's canonical form makes equal records equal bytes.
type Triple struct {
	Pred []byte
	Obj  []byte
}
