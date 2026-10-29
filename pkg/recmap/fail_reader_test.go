//ff:type feature=recmap type=model
//ff:what failReader — a ReaderAt over bytes that fails reads starting at or after a given offset
package recmap

// failReader serves b but errors for any read at off ≥ from.
type failReader struct {
	b    []byte
	from int64
}
