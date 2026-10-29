//ff:type feature=scan type=model
//ff:what Hit — one matching record: its pack offset and a copy of its bytes
package scan

// Hit is a record the Matcher accepted.
type Hit struct {
	Off uint32
	Rec []byte
}
