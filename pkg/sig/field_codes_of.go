//ff:func feature=sig type=decoder control=selection
//ff:what Field.codesOf — the codes of one Object: markers first, then by the field's kind
package sig

func (f *Field) codesOf(o []byte) []uint64 {
	if f.Kind == KindFlag {
		if string(o) == wordNoValue {
			return []uint64{0}
		}
		return []uint64{1}
	}
	switch string(o) {
	case wordNoValue:
		return []uint64{CodeNoValue}
	case wordSomeValue:
		return []uint64{CodeSomeValue}
	}
	if f.Kind == KindBucket {
		return f.bucketCodes(o)
	}
	if e, ok := f.byValue[string(o)]; ok {
		return []uint64{e.Code << (f.Width - e.Len)}
	}
	return []uint64{CodeOther}
}
