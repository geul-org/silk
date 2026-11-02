//ff:func feature=query type=helper control=iteration dimension=1
//ff:what truth — the documents an independent reading of the corpus says hold pred=obj on a person
package query

import (
	"slices"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
)

func truth(t *testing.T, docs []fixture.Doc, pred, obj string) []string {
	t.Helper()
	g := fixture.Grammar(t)
	out := []string{}
	for _, d := range docs {
		if stated(g, d, "P31", "Q5") && stated(g, d, pred, obj) {
			out = append(out, archive.VersionPath("0x00010000", d.ID, d.Version))
		}
	}
	slices.Sort(out)
	return out
}
