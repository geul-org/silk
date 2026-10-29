//ff:func feature=cli type=command control=sequence
//ff:what TestRunIndex — build then verify an archive; missing subcommand, bad flags, unknown subcommand are usage errors
package cli

import (
	"bytes"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

func TestRunIndex(t *testing.T) {
	dir := fixture.Packed(t, nil)
	idx := filepath.Join(t.TempDir(), "idx")
	var out, errb bytes.Buffer
	if runIndex([]string{"build", "-index", idx, dir}, &out, &errb) != 0 {
		t.Fatal(errb.String())
	}
	if runIndex([]string{"verify", "-index", idx, dir}, &out, &errb) != 0 {
		t.Fatal(errb.String())
	}
	if runIndex(nil, &out, &errb) != 2 || runIndex([]string{"build"}, &out, &errb) != 2 || runIndex([]string{"build", "-x", dir}, &out, &errb) != 2 {
		t.Fatal("usage")
	}
	if runIndex([]string{"drop", dir}, &out, &errb) != 2 {
		t.Fatal("unknown")
	}
}
