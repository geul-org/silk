//ff:func feature=query type=scanner control=sequence
//ff:what TestScanRegion — silk patterns, prefixes, and an emit failure
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
	"github.com/geul-org/silk/pkg/sig"
)

func TestScanRegion(t *testing.T) {
	boom := errors.New("stop")
	reg := seg.Region{B: []byte{1, 2}, W: 1}
	n := 0
	count := func([]byte) error { n++; return nil }
	if err := scanRegion(&Plan{Prefix: [][]byte{{1}}}, reg, count); err != nil || n != 1 {
		t.Fatal(err, n)
	}
	if err := scanRegion(&Plan{Pats: []sig.Pattern{{Bits: 8}}}, reg, count); err != nil || n != 3 {
		t.Fatal(err, n)
	}
	if err := scanRegion(&Plan{Pats: []sig.Pattern{{Bits: 8}}}, reg, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if err := scanRegion(&Plan{Prefix: [][]byte{nil}}, reg, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
}
