//ff:func feature=cli type=command control=sequence
//ff:what TestIndexBuild — prints one map line and a summary; a plain tree and an unwritable index fail
package cli

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

func TestIndexBuild(t *testing.T) {
	dir := fixture.Packed(t, nil)
	var out, errb bytes.Buffer
	if indexBuild(dir, filepath.Join(dir, ".silk"), &out, &errb) != 0 || !strings.Contains(out.String(), "0x00010000") || !strings.Contains(errb.String(), "1 maps, 4 documents") {
		t.Fatal(out.String(), errb.String())
	}
	if indexBuild(fixture.Source(), t.TempDir(), &out, &errb) != 1 {
		t.Fatal("plain tree")
	}
	f := filepath.Join(t.TempDir(), "f")
	os.WriteFile(f, nil, 0o644)
	if indexBuild(dir, filepath.Join(f, "idx"), &out, &errb) != 1 {
		t.Fatal("unwritable")
	}
}
