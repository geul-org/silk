//ff:func feature=query type=scanner control=sequence
//ff:what TestBuildLimit — a build under a tiny memory limit, one thread, writes the same index bytes as under a large limit with several
package query

import (
	"context"
	"maps"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func TestBuildLimit(t *testing.T) {
	ctx := context.Background()
	data := fixture.Pack(t, fixture.Corpus(5, 400), nil)
	a, err := archive.Open(ctx, data)
	if err != nil {
		t.Fatal(err)
	}
	make := func(o index.Options) map[string]string {
		dir := filepath.Join(t.TempDir(), "ix")
		declareAll(t, dir)
		if _, err := index.Build(ctx, a, dir, o); err != nil {
			t.Fatal(err)
		}
		return treeHashes(t, dir)
	}
	big := make(index.Options{Limit: 1 << 30, Threads: 4, Policy: index.DefaultPolicy()})
	tiny := make(index.Options{Limit: 2048, Threads: 1, Policy: index.DefaultPolicy()})
	if !maps.Equal(big, tiny) {
		t.Fatalf("index differs:\n%v\n%v", big, tiny)
	}
}
