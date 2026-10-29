//ff:func feature=seg type=helper control=sequence
//ff:what TestInfoRowWidth — partition rows carry the sort slot and 8 bytes, other rows the partition slot too
package seg

import "testing"

func TestInfoRowWidth(t *testing.T) {
	if r, o := (Info{PartWidth: 6, SortWidth: 4}).RowWidth(); r != 12 || o != 18 {
		t.Fatal(r, o)
	}
}
