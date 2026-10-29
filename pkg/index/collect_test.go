//ff:func feature=index type=writer control=iteration dimension=1
//ff:what TestCollect — unreferenced files go; files of the current and pinned generations and paused outputs stay; a pin of a lost manifest keeps nothing
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestCollect(t *testing.T) {
	d := t.TempDir()
	touch := func(rel string) {
		os.MkdirAll(filepath.Dir(filepath.Join(d, rel)), 0o755)
		os.WriteFile(filepath.Join(d, rel), nil, 0o644)
	}
	old := Manifest{Sections: []Section{{Runs: []SegRef{{File: "seg/old"}}}}}
	commit(d, &old)
	cur := Manifest{Generation: 1, Sections: []Section{{Runs: []SegRef{{File: "seg/new"}}}}}
	commit(d, &cur)
	for _, f := range []string{"seg/old", "seg/new", "seg/gone", "seg/out.tmp", "doc/x"} {
		touch(f)
	}
	writeCheckpoints(d, []checkpoint{{Output: "seg/out"}})
	pinGen(d, 1)
	pinGen(d, 9)
	if n := collect(d, cur); n != 2 {
		t.Fatal(n)
	}
	for f, want := range map[string]bool{"seg/old": true, "seg/new": true, "seg/out.tmp": true, "seg/gone": false, "doc/x": false} {
		if _, err := os.Stat(filepath.Join(d, f)); (err == nil) != want {
			t.Fatal(f)
		}
	}
}
