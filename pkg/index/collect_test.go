//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestCollect — files of no generation go; pinned generations stay
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollect(t *testing.T) {
	dir := t.TempDir()
	m := Manifest{Live: "live/a", Grammar: "g", Docs: []DocFile{{File: "doc/a"}}, Paths: Section{Runs: []SegRef{{File: "seg/a"}}}, Sections: []Section{{Decl: "d", Appends: []SegRef{{File: "seg/b"}}}}}
	if err := commit(dir, &m); err != nil {
		t.Fatal(err)
	}
	old := m
	old.Paths.Runs = []SegRef{{File: "seg/old"}}
	if err := commit(dir, &old); err != nil {
		t.Fatal(err)
	}
	for _, f := range []string{"live/a", "doc/a", "seg/a", "seg/b", "seg/old", "seg/stray", "decl/d.json", "grammar/g.json"} {
		os.MkdirAll(filepath.Dir(filepath.Join(dir, f)), 0o755)
		os.WriteFile(filepath.Join(dir, f), nil, 0o644)
	}
	pinGen(dir, 2)
	os.WriteFile(filepath.Join(dir, dirPin, "9-1-1"), nil, 0o644)
	collect(dir, m)
	has := func(f string) bool { _, err := os.Stat(filepath.Join(dir, f)); return err == nil }
	if has("seg/stray") || !has("seg/old") || !has("seg/a") || !has("seg/b") || !has("decl/d.json") {
		t.Fatal("collect")
	}
	keep := map[string]bool{}
	refs(m, keep)
	addSegs(keep, []SegRef{{File: "x"}})
	if !keep["x"] || !keep["doc/a"] {
		t.Fatal(keep)
	}
	es, _ := os.ReadDir(filepath.Join(dir, dirDoc))
	removeUnkept(dir, dirDoc, es, map[string]bool{})
	if has("doc/a") {
		t.Fatal("removeUnkept")
	}
}
