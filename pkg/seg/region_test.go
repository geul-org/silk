//ff:func feature=seg type=scanner control=sequence
//ff:what TestRegion — seek, after and prefix on sorted rows
package seg

import (
	"bytes"
	"testing"
)

func TestRegion(t *testing.T) {
	f := writeSeg(t, "a")
	if f.Doc(f.All().Row(2)) != 3 || testInfo().Doc(row(1, 2, 9)) != 9 {
		t.Fatal("Doc")
	}
	r := f.All()
	if r.Len() != 3 || r.Seek(1, []byte{0, 3}) != 2 || r.After(0, []byte{0, 2}) != 2 {
		t.Fatal(r.Len())
	}
	if p := r.Prefix([]byte{0, 2}); p.Len() != 1 || !bytes.Equal(p.Row(0), row(0, 2, 2)) {
		t.Fatal(p)
	}
	if p := r.Prefix([]byte{9}); p.Len() != 0 {
		t.Fatal(p)
	}
}
