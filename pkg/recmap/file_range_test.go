//ff:func feature=recmap type=reader control=sequence
//ff:what TestFileRange — present keys found; keys between and after ranges absent
package recmap

import "testing"

func TestFileRange(t *testing.T) {
	f := &File{Keys: []KeyRange{{Key: 2, Count: 1}, {Key: 5, Start: 1, Count: 3}}}
	if k, ok := f.Range(5); !ok || k.Count != 3 {
		t.Fatal(k, ok)
	}
	if _, ok := f.Range(3); ok {
		t.Fatal(3)
	}
	if _, ok := f.Range(9); ok {
		t.Fatal(9)
	}
}
