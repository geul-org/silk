//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestRunNoMiss — every document the corpus states a job on (multi-valued, spread-limited) is among the answers; rows with somevalue and other are kept in -keys
package query

import (
	"context"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestRunNoMiss(t *testing.T) {
	ctx := context.Background()
	docs := fixture.Corpus(11, 500)
	_, _, ix := openIndex(t, fixture.Pack(t, docs, nil))
	for field, obj := range map[string]string{"physicist": "Q169470", "writer": "Q36180", "actor": "Q33999", "politician": "Q82955"} {
		want := truth(t, docs, "P106", obj)
		if len(want) == 0 {
			t.Fatal("no truth for", obj)
		}
		q, err := Parse(ix, []string{"entity", "job=" + field})
		if err != nil {
			t.Fatal(err)
		}
		r, err := Run(ctx, ix, q)
		if err != nil {
			t.Fatal(err)
		}
		q.Mode = ModeKeys
		k, err := Run(ctx, ix, q)
		if err != nil {
			t.Fatal(err)
		}
		got, keys := map[string]bool{}, map[string]bool{}
		for _, h := range r.Hits {
			got[ix.Path(h.Doc)] = true
		}
		for _, h := range k.Hits {
			keys[ix.Path(h.Doc)] = true
		}
		if missed := absent(want, got, keys); missed != "" {
			t.Fatal("missed", field, missed)
		}
		if k.Decoded != 0 || len(k.Hits) < len(want) {
			t.Fatal("keys mode", k.Decoded, len(k.Hits))
		}
	}
}
