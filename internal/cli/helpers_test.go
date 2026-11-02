//ff:func feature=cli type=scanner control=iteration dimension=1
//ff:what TestHelpers — sizes, widths, modes and defaults
package cli

import (
	"context"
	"io"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/query"
)

func TestHelpers(t *testing.T) {
	for s, want := range map[string]int64{"12": 12, "2K": 2048, "3M": 3 << 20, "1G": 1 << 30} {
		if n, err := parseSize(s); n != want || err != nil {
			t.Fatal(s, n, err)
		}
	}
	for _, s := range []string{"x", "-1"} {
		if _, err := parseSize(s); err == nil {
			t.Fatal(s)
		}
	}
	if w, err := parseWidths("1,2"); err != nil || len(w) != 2 {
		t.Fatal(w, err)
	}
	if w, err := parseWidths(""); w != nil || err != nil {
		t.Fatal(w, err)
	}
	if _, err := parseWidths("x"); err == nil {
		t.Fatal("bad width")
	}
	f := newFlags("x", io.Discard)
	if f.mode() != query.ModeDocs || f.dir("a") != filepath.Join("a", ".silk") {
		t.Fatal(f.mode())
	}
	f.keys = true
	if f.mode() != query.ModeKeys {
		t.Fatal(f.mode())
	}
	f.keys, f.records = false, true
	if f.mode() != query.ModeRecords {
		t.Fatal(f.mode())
	}
	if p, err := f.parseArgs([]string{"a", "-index", "d", "b"}); err != nil || len(p) != 2 || f.index != "d" {
		t.Fatal(p, err)
	}
	if _, err := f.idlePause(context.Background()); err != nil {
		t.Fatal(err)
	}
	f.idleCPU, f.idleMem = .01, .01
	p, err := f.idlePause(context.Background())
	if err != nil || p() != nil {
		t.Fatal(err)
	}
}
