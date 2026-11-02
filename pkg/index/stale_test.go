//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestStale — evenly spaced probes estimate the dead share
package index

import "testing"

func TestStale(t *testing.T) {
	_, dir, _ := built(t, 4)
	m := current(t, dir)
	var b Bitmap
	for i := range uint32(4) {
		b.Put(i, i%2 == 0)
	}
	s, err := staleShare(dir, m.Sections[0], b)
	if err != nil || s != 0.5 {
		t.Fatal(s, err)
	}
	if s, err := staleShare(dir, Section{}, b); s != 0 || err != nil {
		t.Fatal(s, err)
	}
}
