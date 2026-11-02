//ff:func feature=sig type=scanner control=sequence
//ff:what TestPatternScanEnd — a region past all matches ends without a visit, and a pattern with no room stops at once
package sig

import (
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestPatternScanEnd(t *testing.T) {
	p := Pattern{Bits: 8, Terms: []Term{{Off: 0, Width: 8, Spans: []Span{{1, 1}}}}}
	if err := p.Scan(seg.Region{B: []byte{9}, W: 1}, func([]byte) error { t.Fatal("visited"); return nil }); err != nil {
		t.Fatal(err)
	}
	n := 0
	if err := p.Scan(seg.Region{B: []byte{1, 1}, W: 1}, func([]byte) error { n++; return nil }); err != nil || n != 2 {
		t.Fatal(err, n)
	}
	boom := errors.New("stop")
	if err := p.Scan(seg.Region{B: []byte{1}, W: 1}, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
}
