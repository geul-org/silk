//ff:func feature=grammar type=scanner control=sequence
//ff:what TestGrammarAddBody — a body's DECLARE, BIND and language ENUM land in the grammar; markOdd pairs two low-bit neighbours
package grammar

import (
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func TestGrammarAddBody(t *testing.T) {
	g := &Grammar{Fields: map[uint16][]Field{}, Names: map[string][]uint16{}, Odd: map[uint16]bool{}, Langs: map[string]uint16{}}
	g.addLang("0x0784", "EN")
	g.addLang("zz", "xx")
	if g.Langs["en"] != 0x0784 || len(g.Langs) != 1 {
		t.Fatal(g.Langs)
	}
	id, _ := core.EncodeBind(core.Bind{EdgeID: 2, Subject: core.Record{Kind: core.KindReference, Codec: core.RefLocal << 3, Payload: []core.Word{1}}, Name: []byte("X")})
	dec := core.EncodeDeclare(core.Declare{EdgeID: 1, Target: 0x1050})
	g.addBody([]core.Record{id, dec, {Kind: core.KindMeta}})
	if !slices.Equal(g.Names["X"], []uint16{0x1050}) || len(convertFields([]codecspec.Field{{Name: "a", Codecs: []core.Word{1}}})[0].Codecs) != 1 {
		t.Fatal(g.Names)
	}
	if n := tagNames([]core.Record{id, dec}); n[0x1050] != "X" {
		t.Fatal(n)
	}
	g.Names["P"] = []uint16{0x0845, 0x0844}
	g.markOdd()
	if !g.Odd[0x0845] || g.Odd[0x0844] {
		t.Fatal(g.Odd)
	}
}
