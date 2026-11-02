//ff:func feature=sig type=helper control=sequence
//ff:what Entry.span — the lowest and highest full code under an entry's bits
package sig

func (e *Entry) span(width int) (uint64, uint64) {
	s := width - e.Len
	return e.Code << s, (e.Code+1)<<s - 1
}
