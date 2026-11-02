//ff:func feature=query type=helper control=sequence
//ff:what declareAll — the sections every query test uses
package query

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/index"
)

func declareAll(t *testing.T, dir string) {
	t.Helper()
	g := fixture.Grammar(t)
	must := func(d decl.Decl, err error) {
		t.Helper()
		if err != nil {
			t.Fatal(err)
		}
		if err := index.Declare(dir, d, g); err != nil {
			t.Fatal(err)
		}
	}
	must(decl.NewIndex(g, "triple", "TRIPLE", "predicate", []string{"object"}, []int{6}))
	must(decl.NewIndex(g, "subject", "TRIPLE", "predicate", []string{"subject"}, nil))
	must(decl.NewIndex(g, "label-subj", "LABEL", "object.language", []string{"subject"}, nil))
	must(decl.NewText(g, "label", "LABEL"))
	must(decl.NewText(g, "alias", "ALIAS"))
	must(decl.NewText(g, "bind", "BIND"))
	must(decl.NewSilk("entity", testDef))
	must(decl.NewComposite(g, "label-sem", "LABEL", testSem))
}
