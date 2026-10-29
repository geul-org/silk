//ff:func feature=cli type=helper control=sequence
//ff:what TestLoadSpec — a -spec file wins; update keeps the index's spec; otherwise the default
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/park-jun-woo/silk/pkg/spec"
)

func TestLoadSpec(t *testing.T) {
	sp := spec.Default()
	sp.Tags = sp.Tags[:1]
	p := filepath.Join(t.TempDir(), "s.json")
	os.WriteFile(p, sp.Encode(), 0o644)
	if s, err := loadSpec(p, "", false); err != nil || len(s.Tags) != 1 {
		t.Fatal(err)
	}
	data := fixturePacked(t)
	call("index build -spec " + p + " " + data)
	dir := filepath.Join(data, ".silk")
	if s, err := loadSpec("", dir, true); err != nil || len(s.Tags) != 1 {
		t.Fatal(err)
	}
	if s, _ := loadSpec("", dir, false); len(s.Tags) != 7 {
		t.Fatal("build takes the default")
	}
	if s, _ := loadSpec("", t.TempDir(), true); len(s.Tags) != 7 {
		t.Fatal("no index")
	}
}
