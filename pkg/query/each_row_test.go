//ff:func feature=query type=scanner control=sequence
//ff:what TestEachRow — each row, then an emit failure stops
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestEachRow(t *testing.T) {
	n := 0
	if err := eachRow(seg.Region{B: []byte{1, 2}, W: 1}, func([]byte) error { n++; return nil }); err != nil || n != 2 {
		t.Fatal(err, n)
	}
	boom := errors.New("stop")
	if err := eachRow(seg.Region{B: []byte{1}, W: 1}, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
}
