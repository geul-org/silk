//ff:func feature=index type=scanner control=sequence
//ff:what TestRunTmpBlocked — a tmp path that cannot be made fails the run
package index

import (
	"context"
	"os"
	"path/filepath"
	"testing"
)

func TestRunTmpBlocked(t *testing.T) {
	_, dir, a := setup(t, 1)
	os.WriteFile(filepath.Join(dir, dirTmp), nil, 0o644)
	if _, err := run(context.Background(), a, dir, testOpts, false); err == nil {
		t.Fatal("tmp is a file")
	}
	if _, err := Compact(context.Background(), dir, testOpts, false); err == nil {
		t.Fatal("tmp is a file")
	}
}
