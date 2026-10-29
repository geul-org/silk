//ff:func feature=spec type=encoder control=sequence
//ff:what AppendSlot — a key slot: the leading width bytes of a field record, zero padded
package spec

// AppendSlot appends v cut or padded to width bytes. A record up to width
// bytes keeps all of itself, and since its TAG WORD carries its length, equal
// slots then mean equal records.
func AppendSlot(dst, v []byte, width int) []byte {
	n := min(len(v), width)
	dst = append(dst, v[:n]...)
	return append(dst, make([]byte, width-n)...)
}
