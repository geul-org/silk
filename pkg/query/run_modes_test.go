//ff:func feature=query type=scanner control=sequence
//ff:what TestRunModes — keys without documents, candidates, covered answers, records, the plan text and refusals
package query

import (
	"context"
	"errors"
	"testing"
)

func TestRunModes(t *testing.T) {
	ctx := context.Background()
	_, _, ix := smallIndex(t)
	run := func(mode string, w ...string) (Result, error) {
		q, err := Parse(ix, w)
		if err != nil {
			t.Fatal(err)
		}
		q.Mode = mode
		return Run(ctx, ix, q)
	}
	r, err := run(ModeKeys, "LABEL", "object.language=@en", `object.bytes="Albert E"`)
	if err != nil || r.Decoded != 0 || len(r.Hits) != 1 || !r.Hits[0].Loose || r.Hits[0].QID != 0x40000001 {
		t.Fatal(r, err)
	}
	r, err = run(ModeKeys, "TRIPLE", "predicate=P106", "object=Q98")
	if err != nil || r.Decoded != 0 || len(r.Hits) != 1 || r.Hits[0].Loose {
		t.Fatal(r, err)
	}
	if _, err := run(ModeKeys, "TRIPLE", "object=Q98", "subject=Q2", "predicate=P106", "rank=0"); !errors.Is(err, ErrNotCovered) {
		t.Fatal(err)
	}
	r, err = run(ModeDocs, "LABEL", "object.language=@en", `object.bytes="Short"`)
	if err != nil || len(r.Hits) != 1 {
		t.Fatal("short code equals bcp47", r, err)
	}
	r, err = run(ModeDocs, "LABEL", `object.language=@de`, `object.bytes="Lokal"`)
	if err != nil || len(r.Hits) != 1 {
		t.Fatal("LOCAL object", r, err)
	}
	r, err = run(ModeDocs, "BIND", `name.bytes="albert"`)
	if err != nil || len(r.Hits) != 1 {
		t.Fatal("BIND", r, err)
	}
	r, err = run(ModeDocs, "ALIAS", `object.bytes^="Ein"`)
	if err != nil || len(r.Hits) != 1 || !r.Hits[0].HasOff {
		t.Fatal(r, err)
	}
	r, err = run(ModeDocs, "label-sem", "object.language=@en", `object.bytes^="Albert"`, "type=person")
	if err != nil || len(r.Hits) != 1 {
		t.Fatal(r, err)
	}
	if _, err := run(ModeDocs, "DESCRIPTION", "subject=Q1"); !errors.Is(err, ErrNoPlan) {
		t.Fatal(err)
	}
	r, err = run(ModeDocs, "entity", "job=Q98")
	if err != nil || len(r.Hits) != 1 || !r.Hits[0].Unknown {
		t.Fatal("other type, listed nowhere", r, err)
	}
	r, err = run(ModeDocs, "entity", "sex=male")
	if err != nil || len(r.Hits) != 1 || !r.Hits[0].Unknown {
		t.Fatal("somevalue stays unknown", r, err)
	}
	r, err = run(ModeDocs, "entity", "job=physicist")
	if err != nil || r.Decoded != 0 && len(r.Hits) == 0 {
		t.Fatal(r, err)
	}
}
