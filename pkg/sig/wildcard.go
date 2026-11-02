//ff:func feature=sig type=helper control=sequence
//ff:what wildcard — the code that keeps a field a candidate: `other`, or set for a flag
package sig

func wildcard(f *Field) uint64 {
	if f.Kind == KindFlag {
		return 1
	}
	return CodeOther
}
