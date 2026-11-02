//ff:func feature=index type=builder control=sequence
//ff:what TestBuild — a build forgets what was built and numbers the documents again; a limit is required
package index

import (
	"context"
	"errors"
	"testing"
)

func TestBuild(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	r, err := Build(ctx, a, dir, testOpts)
	if err != nil || r.NewDocs != 2 || r.ReadDocs != 2 {
		t.Fatal(r, err)
	}
	if _, err := Build(ctx, a, dir, Options{}); !errors.Is(err, ErrLimit) {
		t.Fatal(err)
	}
}
