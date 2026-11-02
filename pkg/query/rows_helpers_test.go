//ff:func feature=query type=scanner control=sequence
//ff:what TestRowsHelpers — region reading, row errors, predicate matches
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/seg"
	"github.com/geul-org/silk/pkg/sig"
)

func TestRowsHelpers(t *testing.T) {
	boom := errors.New("stop")
	reg := seg.Region{B: []byte{1, 2}, W: 1}
	if err := eachRow(reg, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if err := scanRegion(&Plan{Prefix: [][]byte{nil}}, reg, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if err := scanRegion(&Plan{Pats: []sig.Pattern{{Bits: 8}}}, reg, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if err := scanRegion(&Plan{Pats: []sig.Pattern{{Bits: 8}}}, reg, func([]byte) error { return nil }); err != nil {
		t.Fatal(err)
	}
	q := Pred{Vals: [][]byte{[]byte("ab")}}
	if !anyEqual(q, []byte("ab")) || anyEqual(q, []byte("abc")) {
		t.Fatal("equal")
	}
	q.Prefix = true
	if !anyEqual(q, []byte("abc")) {
		t.Fatal("prefix")
	}
	if offWanted([]Hit{{HasOff: true, Off: 3}}, grammar.Rec{Off: 4}) || !offWanted([]Hit{{}}, grammar.Rec{Off: 4}) {
		t.Fatal("offWanted")
	}
	if !bitsMatch(&Plan{}, nil) {
		t.Fatal("no bits")
	}
	if exact, loose := rowsMeet([]sig.Pattern{{Bits: 8}}, seg.Region{B: []byte{0, 0, 0, 0, 1}, W: 5}); !exact || !loose {
		t.Fatal(exact, loose)
	}
	if rows(&index.Sec{}) != 0 {
		t.Fatal("rows")
	}
}
