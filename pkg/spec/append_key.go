//ff:func feature=spec type=encoder control=sequence
//ff:what appendKey — one key's slot from a record (nothing when the key has no path)
package spec

// appendKey frames the key's field and appends its slot.
func appendKey(dst, rec []byte, k Key) []byte {
	if len(k.Path) == 0 {
		return dst
	}
	v, _ := Field(rec, k.Path)
	return AppendSlot(dst, v, k.Width)
}
