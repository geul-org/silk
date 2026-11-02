//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestDefEntities — TRIPLEs group by Subject in first-seen order; deprecated, undefined and cut statements drop
package sig

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/grammar"
)

func TestDefEntities(t *testing.T) {
	g := fixture.Grammar(t)
	d := testDef(t)
	s, o := fixture.Tok("Q1"), fixture.Tok("Q2")
	body := []core.Record{
		fixture.Triple(1, fixture.Tok("P31"), s, fixture.Tok("Q5"), 0),
		fixture.Triple(2, fixture.Tok("P106"), s, fixture.Tok("Q36180"), 2),
		fixture.Triple(3, fixture.Tok("P9999"), s, fixture.Tok("Q5"), 0),
		fixture.Triple(4, fixture.Tok("P31"), o, fixture.Tok("Q515"), 1),
		fixture.Label(5, s, "en", "x"),
	}
	var b []byte
	for _, r := range body {
		b = append(b, grammar.RecordOf(r)...)
	}
	doc, err := grammar.Frame(b)
	if err != nil {
		t.Fatal(err)
	}
	es := d.Entities(g, doc)
	if len(es) != 2 || !bytes.Equal(es[0].Subject, s) || len(es[0].Objects) != 1 || len(es[1].Objects) != 1 {
		t.Fatal(es)
	}
	if _, _, _, ok := d.statement(g, b[:6]); ok {
		t.Fatal("cut record")
	}
	ws, _ := core.ReadWords(append(append(fixture.Local(1), fixture.Tok("P31")...), s...))
	if _, _, _, ok := d.statement(g, grammar.RecordOf(core.Record{Kind: core.KindEdge, Codec: 0x0A, Payload: ws})); ok {
		t.Fatal("no Object")
	}
}
