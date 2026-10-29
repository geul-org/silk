//ff:func feature=sig type=builder control=sequence
//ff:what NewTable — an empty signature column for a schema
package sig

// NewTable makes a table of s.Words() words per row.
func NewTable(s Schema) *Table {
	return &Table{Words: s.Words()}
}
