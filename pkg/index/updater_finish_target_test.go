//ff:func feature=index type=scanner control=sequence
//ff:what TestUpdaterFinishTarget — spills become sections; a blocked spill or a section failure stops
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestUpdaterFinishTarget(t *testing.T) {
	_, dir, a := built(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	ts, _ := u.targets()
	var r Report
	if err := u.finishTargets(ts, &r); err != nil {
		t.Fatal(err)
	}
	ts, _ = u.targets()
	ts[0].sp[0].buf.Add(nil, make([]byte, ts[0].sp[0].info.RowWidth))
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	ts[0].sp[0].tmp = file
	if err := u.finishTarget(ts[0], &r); err == nil {
		t.Fatal("spill")
	}
	ts[0].sp[0].tmp = u.tmp
	ts[0].sp[0].buf.Add(nil, make([]byte, ts[0].sp[0].info.RowWidth))
	u.dir = file
	if err := u.finishTarget(ts[0], &r); err == nil {
		t.Fatal("section")
	}
}
