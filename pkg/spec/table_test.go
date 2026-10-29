//ff:func feature=spec type=builder control=sequence
//ff:what TestSpecTable — indexed TAG WORDs map to their tag positions, others to -1
package spec

import "testing"

func TestSpecTable(t *testing.T) {
	tb := Default().Table()
	if len(tb) != 1<<13 || tb[0x100A] != 0 || tb[0x1004] != 2 || tb[0x0000] != -1 || tb[0x1FFF] != -1 {
		t.Fatal(tb[0x100A], tb[0x1004])
	}
}
