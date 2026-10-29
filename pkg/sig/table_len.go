//ff:func feature=sig type=helper control=sequence
//ff:what Table.Len — the number of rows
package sig

// Len is len(Data)/Words (0 for a table without a width).
func (t *Table) Len() int {
	if t.Words <= 0 {
		return 0
	}
	return len(t.Data) / t.Words
}
