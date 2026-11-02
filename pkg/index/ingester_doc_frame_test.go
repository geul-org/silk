//ff:func feature=index type=scanner control=sequence
//ff:what TestIngesterDocFrame — a document that does not frame stops with ErrFrame
package index

import (
	"context"
	"errors"
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
)

func TestIngesterDocFrame(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	g := &ingester{u: u}
	docs, _ := openDocs(dir, &u.m)
	defer docs.Close()
	e := docs.At(0)
	e.Len = 2
	if err := g.doc(ctx, job{entry: e}, nil); !errors.Is(err, grammar.ErrFrame) {
		t.Fatal(err)
	}
}
