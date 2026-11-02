//ff:func feature=sig type=helper control=sequence
//ff:what sortRows — sorted, byte-equal rows dropped
package sig

import (
	"bytes"
	"slices"
)

func sortRows(rows [][]byte) [][]byte {
	slices.SortFunc(rows, bytes.Compare)
	return slices.CompactFunc(rows, bytes.Equal)
}
