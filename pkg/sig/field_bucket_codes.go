//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what Field.bucketCodes — the buckets a date or number's span crosses, one code each; CodeOther when it cannot be read or no bucket holds it
package sig

func (f *Field) bucketCodes(o []byte) []uint64 {
	lo, hi, ok := yearSpan(o)
	if !ok {
		return []uint64{CodeOther}
	}
	var out []uint64
	for _, e := range f.Codes {
		if e.Lo <= hi && lo < e.Hi {
			out = append(out, e.Code)
		}
	}
	if out == nil {
		return []uint64{CodeOther}
	}
	return out
}
