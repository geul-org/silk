//ff:func feature=cli type=command control=sequence
//ff:what TestIndexVerify — a built index passes; a plain tree and a missing index fail
package cli

import (
	"bytes"
	"path/filepath"
	"strings"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

func TestIndexVerify(t *testing.T) {
	dir := fixture.Packed(t, nil)
	idx := filepath.Join(dir, ".silk")
	var out, errb bytes.Buffer
	indexBuild(dir, idx, &out, &errb)
	out.Reset()
	if indexVerify(dir, idx, &out, &errb) != 0 || !strings.Contains(out.String(), "1 maps ok") {
		t.Fatal(out.String(), errb.String())
	}
	if indexVerify(fixture.Source(), idx, &out, &errb) != 1 {
		t.Fatal("plain tree")
	}
	if indexVerify(dir, t.TempDir(), &out, &errb) != 1 {
		t.Fatal("no index")
	}
}
