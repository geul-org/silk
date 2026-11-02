//ff:func feature=query type=helper control=iteration dimension=1
//ff:what answer — a query's hits as sorted lines: path, offset when known, candidate marks
package query

import (
	"context"
	"fmt"
	"slices"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func answer(t *testing.T, ix *index.Index, scan bool, words ...string) []string {
	t.Helper()
	q, err := Parse(ix, words)
	if err != nil {
		t.Fatal(err)
	}
	q.Scan = scan
	r, err := Run(context.Background(), ix, q)
	if err != nil {
		t.Fatal(words, err)
	}
	out := []string{}
	for _, h := range r.Hits {
		out = append(out, fmt.Sprintf("%s +%d", ix.Path(h.Doc), h.Off))
	}
	slices.Sort(out)
	return slices.Compact(out)
}
