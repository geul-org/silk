//ff:func feature=recmap type=reader control=iteration dimension=1
//ff:what TestFileDoc — offsets inside a document find it; before the first, in a gap and past the last find none
package recmap

import "testing"

func TestFileDoc(t *testing.T) {
	f := &File{Docs: []DocEntry{{Off: 64, Len: 10, ID: 1}, {Off: 80, Len: 4, ID: 2}}}
	if d, ok := f.Doc(73); !ok || d.ID != 1 {
		t.Fatal(d, ok)
	}
	if d, ok := f.Doc(80); !ok || d.ID != 2 {
		t.Fatal(d, ok)
	}
	for _, off := range []uint32{10, 74, 84} {
		if _, ok := f.Doc(off); ok {
			t.Fatal(off)
		}
	}
}
