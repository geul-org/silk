//ff:func feature=index type=scanner control=sequence
//ff:what TestIndexSameSec — equal sections pass, different rows do not
package index

import (
	"context"
	"errors"
	"testing"
)

func TestIndexSameSec(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	ix, _ := Open(ctx, a, dir, 0)
	defer ix.Close()
	if err := ix.sameSec(ix, "t", ix.secs["t"]); err != nil {
		t.Fatal(err)
	}
	o := Index{Live: Bitmap{}, secs: ix.secs}
	if err := ix.sameSec(&o, "t", ix.secs["t"]); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
}
