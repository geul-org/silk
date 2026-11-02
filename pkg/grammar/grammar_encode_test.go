//ff:func feature=grammar type=scanner control=iteration dimension=1
//ff:what TestGrammarEncode — every token form at a record field and a raw field
package grammar

import (
	"bytes"
	"errors"
	"testing"
)

func TestGrammarEncode(t *testing.T) {
	g := testGrammar(t)
	rec, _ := g.Resolve(0x1004, "object")
	raw, _ := g.Resolve(0x100A, "rank")
	rest, _ := g.Resolve(0x1004, "object.bytes")
	lang, _ := g.Resolve(0x1004, "object.language")
	if v, err := g.Encode(rec, "Q5"); err != nil || !bytes.Equal(v[0], tok("Q5")) {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(rec, "0x00a00005"); err != nil || !bytes.Equal(v[0], tok("Q5")) {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(lang, "@ko"); err != nil || len(v) != 2 {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(rec, `"김치"@ko`); err != nil || len(v) != 2 {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(rec, `"김치"@@ko`); err != nil || len(v) != 1 {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(rec, `"abc"`); err != nil || grammarTarget(v[0]) != TagTextOdd {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(rest, `"김치"`); err != nil || string(v[0]) != "김치" {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(rest, "0x4142"); err != nil || string(v[0]) != "AB" {
		t.Fatal(v, err)
	}
	if v, err := g.Encode(raw, "1"); err != nil || !bytes.Equal(v[0], []byte{0, 1}) {
		t.Fatal(v, err)
	}
	if v, err := g.encodeRaw(Leaf{Field: Field{Width: 8}}, "1"); err != nil || len(v[0]) != 16 {
		t.Fatal(v, err)
	}
	bad := []struct {
		l Leaf
		s string
	}{{rec, ""}, {rec, "0xZZ"}, {rec, "nope"}, {rest, "12"}, {rest, `"open`}, {rest, "0xZ"}, {raw, "x"}, {rec, `"x"@@zz-nope`}, {rec, `"x`}}
	for _, c := range bad {
		if _, err := g.Encode(c.l, c.s); !errors.Is(err, ErrSyntax) {
			t.Fatal(c.s, err)
		}
	}
	if _, err := g.langRecords("qqq", true); !errors.Is(err, ErrSyntax) {
		t.Fatal(err)
	}
	if rs, err := g.langRecords("qqq", false); err != nil || len(rs) != 1 {
		t.Fatal(rs, err)
	}
	if wrapHex(nil) != nil {
		t.Fatal("nil")
	}
	if len(recordsOf(nil)) != 0 {
		t.Fatal("empty")
	}
}
