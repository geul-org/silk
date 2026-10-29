//ff:func feature=index type=writer control=sequence
//ff:what TestSaveRound — commits when merges were applied, pauses when checkpoints remain; failing checkpoint or manifest writes fail
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestSaveRound(t *testing.T) {
	d := t.TempDir()
	m := Manifest{}
	var rep Report
	if err := saveRound(d, &m, nil, 1, &rep); err != nil || m.Generation != 1 || rep.Generation != 1 {
		t.Fatal(err)
	}
	if err := saveRound(d, &m, []checkpoint{{Section: "S"}}, 0, &rep); !errors.Is(err, ErrPaused) || m.Generation != 1 || !rep.Paused {
		t.Fatal(err)
	}
	os.Mkdir(filepath.Join(d, checkpointFile+".tmp"), 0o755)
	if err := saveRound(d, &m, []checkpoint{{}}, 0, &rep); err == nil || errors.Is(err, ErrPaused) {
		t.Fatal(err)
	}
	block(d, dirManifest)
	if err := saveRound(d, &m, nil, 1, &rep); err == nil {
		t.Fatal("blocked manifest directory")
	}
}
