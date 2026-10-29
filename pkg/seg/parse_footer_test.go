//ff:func feature=seg type=decoder control=iteration dimension=1
//ff:what TestParseFooter — a good footer parses; wrong length, magic, version, flags, reserved bytes, widths or padding fail
package seg

import (
	"errors"
	"testing"
)

func TestParseFooter(t *testing.T) {
	good := encodeFooter(Info{PartWidth: 2, SortWidth: 2})
	if _, err := parseFooter(good); err != nil {
		t.Fatal(err)
	}
	edits := []func(b []byte) []byte{
		func(b []byte) []byte { return b[:100] },
		func(b []byte) []byte { b[0] = 'X'; return b },
		func(b []byte) []byte { b[127] = 'X'; return b },
		func(b []byte) []byte { b[9] = 2; return b },
		func(b []byte) []byte { b[11] = 4; return b },
		func(b []byte) []byte { b[19] = 1; return b },
		func(b []byte) []byte { b[14] = 1; return b },
		func(b []byte) []byte { b[16] = 1; return b },
		func(b []byte) []byte { b[110] = 1; return b },
	}
	for i, e := range edits {
		if _, err := parseFooter(e(append([]byte{}, good...))); !errors.Is(err, ErrFormat) {
			t.Fatalf("edit %d: %v", i, err)
		}
	}
}
