//ff:func feature=query type=scanner control=sequence
//ff:what TestRunnerSelectRows — covered, keys, silk and text paths through one segment
package query

import "testing"

func TestRunnerSelectRows(t *testing.T) {
	_, _, ix := smallIndex(t)
	r := testRunner(ix)
	run := func(mode string, w ...string) []Hit {
		q, err := Parse(ix, w)
		if err != nil {
			t.Fatal(err)
		}
		p, err := makePlan(ix, &q.Conds[0], "")
		if err != nil {
			t.Fatal(err)
		}
		hs, err := r.selectRows(p, Query{Mode: mode})
		if err != nil {
			t.Fatal(err)
		}
		return hs
	}
	if h := run(ModeDocs, "TRIPLE", "predicate=P106", "object=Q98"); len(h) != 1 || !h[0].HasOff {
		t.Fatal(h)
	}
	if h := run(ModeKeys, "entity", "job=physicist"); len(h) == 0 {
		t.Fatal(h)
	}
	if h := run(ModeDocs, "entity", "job=physicist"); len(h) != 1 {
		t.Fatal(h)
	}
	if h := run(ModeDocs, "LABEL", `object.bytes="Short"`); len(h) != 1 {
		t.Fatal(h)
	}
}
