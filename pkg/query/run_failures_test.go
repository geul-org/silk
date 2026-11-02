//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestRunFailures — a pack that went away or a cancelled read reaches the caller from every stage
package query

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/sig"
)

func TestRunFailures(t *testing.T) {
	ctx := context.Background()
	_, _, ix := smallIndex(t)
	queries := [][]string{
		{"LABEL", `object.bytes="Short"`},
		{"entity", "job=Q98"},
		{"TRIPLE", "predicate=P106", "object=Q98", "entity", "job=Q98"},
	}
	var parsed []Query
	for _, w := range queries {
		q, err := Parse(ix, w)
		if err != nil {
			t.Fatal(err)
		}
		parsed = append(parsed, q)
	}
	keep := ix.Manifest.Packs[0].Hash
	ix.Manifest.Packs[0].Hash = "gone"
	defer func() { ix.Manifest.Packs[0].Hash = keep }()
	for i, q := range parsed {
		if _, err := Run(ctx, ix, q); !errors.Is(err, index.ErrStale) {
			t.Fatal(i, err)
		}
		q.Scan = true
		if _, err := Run(ctx, ix, q); !errors.Is(err, index.ErrStale) {
			t.Fatal(i, err)
		}
	}
	r := &runner{ix: ix, ctx: ctx, seen: map[uint32]bool{}}
	if _, _, err := r.holds(&Plan{}, 0); !errors.Is(err, index.ErrStale) {
		t.Fatal(err)
	}
	if _, err := r.silkFilter(&Plan{Name: "none"}, nil, Query{}); !errors.Is(err, ErrNoPlan) {
		t.Fatal(err)
	}
	ix.Manifest.Packs[0].Hash = keep
	bad := &Cond{Decl: &decl.Decl{Silk: "x"}, Silk: []sig.Cond{{}}}
	if _, err := r.scanSilk(0, &grammar.Doc{}, bad); err == nil {
		t.Fatal("bad silk")
	}
	d, _ := grammar.Frame(nil)
	c := &Cond{Decl: &decl.Decl{Silk: testDef}, Silk: []sig.Cond{{Field: "nope"}}}
	if ok, err := r.scanSilk(0, d, c); ok || err != nil {
		t.Fatal(ok, err)
	}
	if hs, err := r.scanCond(0, d, c); hs != nil || err != nil {
		t.Fatal(hs, err)
	}
	if hs, err := r.scanCond(0, d, bad); hs != nil || err == nil {
		t.Fatal(hs, err)
	}
	def, _ := sig.Parse(testDef)
	b, _ := ix.Bytes(ctx, 0)
	doc, _ := grammar.Frame(b)
	cs := []sig.Cond{{Field: "nope", Value: "1"}}
	if ok, err := r.scanSilk(0, doc, &Cond{Decl: &decl.Decl{Silk: testDef}, Silk: cs}); ok || err == nil {
		t.Fatal(ok, err)
	}
	if _, _, err := r.holds(&Plan{Def: def, Cond: &Cond{Silk: cs}}, 0); err == nil {
		t.Fatal("bad field")
	}
	if subjectQID(ix, &grammar.Doc{B: []byte{0x10, 0x03}}, grammar.Rec{Len: 2, Tag: 0x1003}) != 0 {
		t.Fatal("no Subject")
	}
}
