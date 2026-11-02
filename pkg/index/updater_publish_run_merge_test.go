//ff:func feature=index type=scanner control=sequence
//ff:what TestUpdaterPublishRunMerge — a run merged with appends drops dead rows; a blocked merge fails
package index

import (
	"encoding/binary"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestUpdaterPublishRunMerge(t *testing.T) {
	_, dir, a := built(t, 1)
	u := &updater{ar: a, dir: dir, o: testOpts, m: current(t, dir), tmp: t.TempDir()}
	u.loadLive()
	mk := func(name string, seq uint32) string {
		p := filepath.Join(u.tmp, name)
		w, _ := seg.NewWriter(p, pathInfo())
		w.Add(nil, binary.BigEndian.AppendUint32(make([]byte, pathKey), seq))
		w.Finish()
		return p
	}
	ap := mk("append", 0)
	os.Rename(ap, filepath.Join(dir, "seg", "append"))
	s := Section{Appends: []SegRef{{File: "seg/append", Rows: 1}}}
	if err := u.publishRun(&s, mk("new", 0), 1); err != nil || len(s.Appends) != 0 || s.Runs[0].Rows != 2 {
		t.Fatal(err, s)
	}
	q := Section{Appends: []SegRef{{File: "seg/none"}}}
	if err := u.publishRun(&q, mk("new2", 0), 1); err == nil {
		t.Fatal("missing append")
	}
	u.tmp = filepath.Join(dir, currentFile)
	if err := u.publishRun(&Section{Appends: []SegRef{{File: s.Runs[0].File}}}, filepath.Join(dir, s.Runs[0].File), 1); err == nil {
		t.Fatal("tmp under a file")
	}
}
