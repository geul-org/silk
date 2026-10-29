//ff:func feature=seg type=helper control=sequence
//ff:what TestSortRows — rows ordered by key, then document, then offset
package seg

import (
	"bytes"
	"testing"
)

func TestSortRows(t *testing.T) {
	in := bytes.Join([][]byte{row("b1", 1, 0), row("a1", 2, 0), row("a1", 1, 5), row("a1", 1, 2)}, nil)
	want := bytes.Join([][]byte{row("a1", 1, 2), row("a1", 1, 5), row("a1", 2, 0), row("b1", 1, 0)}, nil)
	if got := SortRows(in, 10); !bytes.Equal(got, want) {
		t.Fatalf("%x", got)
	}
}
