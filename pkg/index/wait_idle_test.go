//ff:func feature=index type=helper control=sequence
//ff:what TestWaitIdle — nil without a hook, the hook's error with one
package index

import (
	"context"
	"io"
	"testing"
)

func TestWaitIdle(t *testing.T) {
	ctx := context.Background()
	if waitIdle(ctx, Options{}) != nil {
		t.Fatal("no hook")
	}
	if waitIdle(ctx, Options{Pause: func(context.Context) error { return io.EOF }}) != io.EOF {
		t.Fatal("hook")
	}
}
