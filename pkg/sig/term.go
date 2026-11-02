//ff:type feature=sig type=model
//ff:what Term — what one field of a row may hold: Spans ascending; Exact are the codes that answer without a document; Truth the codes that hold when read from statements
package sig

type Term struct {
	Off, Width int
	Spans      []Span
	Exact      []Span
	Truth      []Span
}
