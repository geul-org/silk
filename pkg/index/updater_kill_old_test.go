//ff:func feature=index type=scanner control=sequence
//ff:what TestUpdaterKillOld — every number an older segment holds for a path dies
package index

import (
	"encoding/binary"
	"path/filepath"
	"slices"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestUpdaterKillOld(t *testing.T) {
	p := filepath.Join(t.TempDir(), "o")
	w, _ := seg.NewWriter(p, pathInfo())
	k := identity(DocEntry{ID: 1})
	w.Add(nil, binary.BigEndian.AppendUint32(slices.Clone(k), 0))
	w.Add(nil, binary.BigEndian.AppendUint32(slices.Clone(k), 1))
	w.Finish()
	f, _ := seg.Open(p)
	defer f.Close()
	u := &updater{}
	u.live.Put(0, true)
	u.live.Put(1, true)
	u.live.Put(2, true)
	u.killOld(k, []*seg.File{f})
	if u.live.Get(0) || u.live.Get(1) || !u.live.Get(2) {
		t.Fatal(u.live)
	}
}
