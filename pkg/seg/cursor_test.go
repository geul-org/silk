//ff:func feature=seg type=scanner control=sequence
//ff:what TestCursor — load and next walk every partition; the heap orders by key then row
package seg

import (
	"bytes"
	"testing"
)

func TestCursor(t *testing.T) {
	a := writeSeg(t, "a")
	c := &cursor{f: a}
	if !c.load() || !bytes.Equal(c.row(), row(0, 1, 1)) {
		t.Fatal(c)
	}
	if !c.next() || !c.next() || string(c.key) != "pb" || c.next() || c.load() {
		t.Fatal(c.key)
	}
	d := &cursor{f: a}
	e := &cursor{f: a, part: 1}
	d.load()
	e.load()
	h := curHeap{e, d}
	if h.Len() != 2 || h.Less(0, 1) || !h.Less(1, 0) {
		t.Fatal(h)
	}
	d2 := &cursor{f: a}
	d2.load()
	d2.next()
	k := curHeap{d, d2}
	if !k.Less(0, 1) {
		t.Fatal("row order")
	}
	h.Swap(0, 1)
	h.Push(d2)
	if h.Pop() != d2 || len(h) != 2 {
		t.Fatal(h)
	}
	hp := newHeap([]*File{a})
	advance(hp, (*hp)[0])
	if hp.Len() != 1 {
		t.Fatal(hp.Len())
	}
	(*hp)[0].i, (*hp)[0].part = 2, a.Parts
	advance(hp, (*hp)[0])
	if hp.Len() != 0 {
		t.Fatal(hp.Len())
	}
}
