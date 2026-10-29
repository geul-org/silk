//ff:func feature=seg type=encoder control=iteration dimension=1
//ff:what TestEncodeFooter — 128 bytes, magic at both ends, round-trips through parseFooter (sorted and not)
package seg

import "testing"

func TestEncodeFooter(t *testing.T) {
	for _, s := range []bool{true, false} {
		i := Info{Tag: 0x100A, PartWidth: 6, SortWidth: 6, Sorted: s, Parts: 3, Rows: 9, Other: 2, Spec: [32]byte{1}, Sum: [32]byte{2}}
		b := encodeFooter(i)
		got, err := parseFooter(b)
		if len(b) != footerSize || string(b[:8]) != magic || err != nil || got != i {
			t.Fatal(err, got)
		}
	}
}
