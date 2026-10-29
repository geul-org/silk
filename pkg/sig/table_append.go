//ff:func feature=sig type=builder control=sequence
//ff:what Table.Append — add one signature row; its length must be Words
package sig

import "fmt"

// Append adds s as the next row.
func (t *Table) Append(s []uint64) error {
	if len(s) != t.Words {
		return fmt.Errorf("%w: signature of %d words, table has %d", ErrValue, len(s), t.Words)
	}
	t.Data = append(t.Data, s...)
	return nil
}
