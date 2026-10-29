//ff:type feature=sig type=model
//ff:what Schema — a signature declaration: total width in bits (a power of two, at least 16) and its fields
package sig

// Schema declares a signature layout. Validate before use.
type Schema struct {
	Width  int
	Fields []Field
}
