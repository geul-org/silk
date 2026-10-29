//ff:func feature=index type=builder control=sequence
//ff:what TestResumeTasks — a checkpoint whose inputs are still in its section becomes a task (file counter moved past its output); a stale one loses its temporary file
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestResumeTasks(t *testing.T) {
	d := t.TempDir()
	m := Manifest{NextFile: 1, Sections: []Section{{Name: "S", Runs: []SegRef{{File: "seg/a", Rows: 4}}}}}
	stale := filepath.Join(d, "seg", "00000000000000f0.sseg.tmp")
	os.MkdirAll(filepath.Dir(stale), 0o755)
	os.WriteFile(stale, nil, 0o644)
	cs := []checkpoint{
		{Section: "S", Kind: TaskMajor, Inputs: []string{"seg/a"}, Output: "seg/00000000000000ff.sseg"},
		{Section: "S", Inputs: []string{"seg/gone"}, Output: "seg/00000000000000f0.sseg"},
	}
	cs[0].State.Min = 3
	ts, by := resumeTasks(d, &m, cs)
	if len(ts) != 1 || ts[0].Inputs[0].Rows != 4 || ts[0].Min != 3 || m.NextFile != 0xff || by["seg/00000000000000ff.sseg"] == nil {
		t.Fatal(ts, m.NextFile)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatal("stale output kept")
	}
}
