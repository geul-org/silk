//ff:func feature=query type=scanner control=sequence
//ff:what TestAnyExact — some pattern exact, some pattern matching, none
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/sig"
)

func TestAnyExact(t *testing.T) {
	p := sig.Pattern{Bits: 8, Terms: []sig.Term{{Off: 0, Width: 8, Spans: []sig.Span{{Lo: 1, Hi: 2}}, Exact: []sig.Span{{Lo: 1, Hi: 1}}}}}
	ps := []sig.Pattern{p}
	if !anyExact(ps, []byte{1}) || anyExact(ps, []byte{2}) {
		t.Fatal("exact")
	}
	if !anyMatch(ps, []byte{2}) || anyMatch(ps, []byte{3}) {
		t.Fatal("match")
	}
}
