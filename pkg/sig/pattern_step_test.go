//ff:func feature=sig type=scanner control=sequence
//ff:what TestPatternStep — a match moves one, a miss jumps, the end stops, the callback's error returns
package sig

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestPatternStep(t *testing.T) {
	p := Pattern{Bits: 8, Terms: []Term{{Off: 0, Width: 8, Spans: []Span{{5, 5}}}}}
	r := seg.Region{B: []byte{1, 5, 9}, W: 1}
	key := make([]byte, 1)
	visit := func([]byte) error { return nil }
	if n, err := p.step(r, 0, key, visit); n != 1 || err != nil {
		t.Fatal(n, err)
	}
	if n, err := p.step(r, 1, key, visit); n != 2 || err != nil {
		t.Fatal(n, err)
	}
	if n, err := p.step(r, 2, key, visit); n != -1 || err != nil {
		t.Fatal(n, err)
	}
	boom := errors.New("stop")
	if _, err := p.step(r, 1, key, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
}
