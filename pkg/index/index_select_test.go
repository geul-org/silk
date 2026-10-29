//ff:func feature=index type=scanner control=sequence
//ff:what TestIndexSelect — live rows of every segment of a section are emitted, dead ones counted; an emit error stops
package index

import (
	"errors"
	"io"
	"testing"

	"github.com/park-jun-woo/silk/pkg/seg"
)

func TestIndexSelect(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	all := func(uint32, uint32) error { return nil }
	st, err := ix.Select(2, seg.Sel{}, all)
	if err != nil || st.Segments != 1 || st.Hits == 0 || st.Dead != 0 {
		t.Fatal(st, err)
	}
	ix.Live.Put(0, false)
	st2, _ := ix.Select(2, seg.Sel{}, all)
	if st2.Dead == 0 || st2.Hits+st2.Dead != st.Hits {
		t.Fatal(st2)
	}
	if _, err := ix.Select(2, seg.Sel{}, func(uint32, uint32) error { return io.EOF }); !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
}
