//ff:func feature=decl type=helper control=iteration dimension=1
//ff:what testDoc — a framed document on Q1: two jobs, a LABEL, an ALIAS twice with one text head, a BIND, a LABEL through a LOCAL
package decl

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func testDoc(t *testing.T) *grammar.Doc {
	t.Helper()
	s := fixture.Tok("Q1")
	recs := []core.Record{
		fixture.Triple(1, fixture.Tok("P31"), s, fixture.Tok("Q5"), 0),
		fixture.Triple(2, fixture.Tok("P106"), s, fixture.Tok("Q1"), 0),
		fixture.Triple(3, fixture.Tok("P106"), s, fixture.Tok("Q2"), 0),
		fixture.Label(4, s, "en", "Hello world"),
		fixture.Alias(5, s, "en", "Hello world two"),
		fixture.Alias(6, s, "en", "Hello world three"),
		fixture.Bind(7, s, "main"),
		fixture.Label(8, []byte{0x10, 0x00}, "de", "no subject"),
	}
	recs = append(recs, fixture.LabelLocal(9, 50, fixture.Tok("Q2"), "ko", "김치")...)
	var b []byte
	for _, r := range recs {
		b = append(b, grammar.RecordOf(r)...)
	}
	d, err := grammar.Frame(b)
	if err != nil {
		t.Fatal(err)
	}
	return d
}
