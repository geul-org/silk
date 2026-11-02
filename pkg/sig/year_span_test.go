//ff:func feature=sig type=scanner control=sequence
//ff:what TestYearSpan — each date and number form, and what cannot be read
package sig

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestYearSpan(t *testing.T) {
	ws, _ := core.ReadWords([]byte{0, 0, 0x0A, 0xD6})
	g := grammar.RecordOf(core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: ws})
	if a, b, ok := yearSpan(g); !ok || a != 1977 || b != 1977 {
		t.Fatal(a, b, ok)
	}
	if _, _, ok := yearSpan([]byte{0}); ok {
		t.Fatal("short")
	}
	if _, _, ok := yearSpan(fixture.Tok("Q1")); ok {
		t.Fatal("reference")
	}
	if _, _, ok := yearSpan(append(slices.Clone(g), 0, 0)); ok {
		t.Fatal("wide date")
	}
	if a, b, ok := yearSpan(edtf("195X-01")); !ok || a != 1950 || b != 1959 {
		t.Fatal(a, b, ok)
	}
	if a, b, ok := edtfYears("-004X"); !ok || a != -49 || b != -40 {
		t.Fatal(a, b, ok)
	}
	gd := grammar.RecordOf(core.EncodeBytesValue(core.CodecGEDCOMDateEven, []byte("JULIAN 1 JAN 1500")))
	if a, _, ok := yearSpan(gd); !ok || a != 1500 {
		t.Fatal(a, ok)
	}
	if a, _, ok := gedcomYears("JULIAN 44 BCE"); !ok || a != -44 {
		t.Fatal(a, ok)
	}
	if _, _, ok := gedcomYears(""); ok {
		t.Fatal("empty")
	}
	if _, _, ok := gedcomYears("ABC"); ok {
		t.Fatal("no number")
	}
	u := grammar.RecordOf(core.EncodeUintValue(1999))
	if a, _, ok := yearSpan(u); !ok || a != 1999 {
		t.Fatal(a, ok)
	}
	if a, _, ok := intYears([]byte{0xFF, 0xFF}, true); !ok || a != -1 {
		t.Fatal(a, ok)
	}
	if _, _, ok := intYears(nil, false); ok {
		t.Fatal("empty")
	}
	if _, _, ok := intYears(make([]byte, 9), false); ok {
		t.Fatal("wide")
	}
}
