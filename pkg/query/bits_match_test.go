//ff:func feature=query type=scanner control=sequence
//ff:what TestBitsMatch — a composite row's mask holds or fails
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/sig"
)

func TestBitsMatch(t *testing.T) {
	p := sig.Pattern{Bits: 32, Terms: []sig.Term{{Off: 0, Width: 8, Spans: []sig.Span{{Lo: 1, Hi: 1}}}}}
	row := make([]byte, 20)
	row[decl.TextWidth] = 1
	if !bitsMatch(&Plan{Bits: []sig.Pattern{p}}, row) {
		t.Fatal("holds")
	}
	row[decl.TextWidth] = 2
	if bitsMatch(&Plan{Bits: []sig.Pattern{p}}, row) {
		t.Fatal("fails")
	}
}
