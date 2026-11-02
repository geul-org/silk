//ff:func feature=query type=scanner control=sequence
//ff:what TestSubjectQID — a Subject's QID word, or 0 without a Subject
package query

import (
	"slices"
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
)

func TestSubjectQID(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	d, _ := r.doc(1)
	i := slices.IndexFunc(d.Recs, func(r grammar.Rec) bool { return r.Tag == 0x1004 })
	if n := subjectQID(ix, d, d.Recs[i]); n != 0x40000002 {
		t.Fatalf("%x", n)
	}
	y := &grammar.Doc{B: []byte{0x10, 0x03}}
	if n := subjectQID(ix, y, grammar.Rec{Len: 2}); n != 0 {
		t.Fatal(n)
	}
}
