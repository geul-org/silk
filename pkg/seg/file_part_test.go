//ff:func feature=seg type=scanner control=sequence
//ff:what TestFilePart — table lookups, partition order, closing
package seg

import "testing"

func TestFilePart(t *testing.T) {
	f := writeSeg(t, "a")
	k, r := f.PartAt(1)
	if string(k) != "pb" || r.Len() != 1 {
		t.Fatal(k, r)
	}
	if k, s, n := f.entry(0); string(k) != "pa" || s != 0 || n != 2 {
		t.Fatal(k, s, n)
	}
	if r, ok := f.Part([]byte("pa")); !ok || r.Len() != 2 {
		t.Fatal(r, ok)
	}
	if r, ok := f.Part([]byte("pc")); ok || r.Len() != 0 {
		t.Fatal(r, ok)
	}
	if r, ok := f.Part([]byte("p0")); ok || r.W != 6 {
		t.Fatal(r, ok)
	}
}
