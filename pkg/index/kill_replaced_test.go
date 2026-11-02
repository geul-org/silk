//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestKillReplaced — in one path group the last number lives; older segments' numbers die
package index

import (
	"encoding/binary"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestKillReplaced(t *testing.T) {
	dir := t.TempDir()
	mk := func(name string, rows ...[]byte) *seg.File {
		p := filepath.Join(dir, name)
		w, _ := seg.NewWriter(p, pathInfo())
		for _, r := range rows {
			w.Add(nil, r)
		}
		w.Finish()
		f, _ := seg.Open(p)
		t.Cleanup(func() { f.Close() })
		return f
	}
	key := func(id uint64, seq uint32) []byte {
		return binary.BigEndian.AppendUint32(identity(DocEntry{ID: id, Version: 1}), seq)
	}
	u := &updater{}
	for i := range uint32(6) {
		u.live.Put(i, true)
	}
	old := mk("old", key(1, 0), key(2, 1))
	nw := mk("new", key(1, 2), key(1, 3), key(3, 4))
	u.killReplaced(nw, []*seg.File{old})
	if u.live.Get(0) || !u.live.Get(1) || u.live.Get(2) || !u.live.Get(3) || !u.live.Get(4) {
		t.Fatal(u.live)
	}
}
