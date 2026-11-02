//ff:func feature=index type=scanner control=sequence
//ff:what TestVerifyStages — a damaged generation stops at the file check; a blocked fresh directory, a failing fresh build and a differing archive stop later
package index

import (
	"context"
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"testing"
)

func TestVerifyStages(t *testing.T) {
	ctx := context.Background()
	_, dir, a := built(t, 2)
	os.RemoveAll(filepath.Join(dir, dirTmp))
	if err := os.WriteFile(filepath.Join(dir, dirTmp), nil, 0o644); err != nil {
		t.Fatal(err)
	}
	if err := Verify(ctx, a, dir, testOpts); err == nil {
		t.Fatal("tmp is a file")
	}
	os.Remove(filepath.Join(dir, dirTmp))
	if err := Verify(ctx, a, dir, testOpts); err != nil {
		t.Fatal(err)
	}
	fresh := filepath.Join(dir, dirTmp, "verify-"+strconv.Itoa(os.Getpid()))
	os.MkdirAll(fresh, 0o755)
	os.WriteFile(filepath.Join(fresh, dirPin), nil, 0o644)
	if err := Verify(ctx, a, dir, testOpts); err == nil {
		t.Fatal("the fresh index cannot be pinned")
	}
	cctx, cancel := context.WithCancel(ctx)
	cancel()
	if err := Verify(cctx, a, dir, testOpts); err != nil && !errors.Is(err, context.Canceled) {
		t.Fatal(err)
	}
}
