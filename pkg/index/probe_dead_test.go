//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestProbeDead — every row probed in a short segment, spaced in a long one
package index

import (
	"encoding/binary"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestProbeDead(t *testing.T) {
	dir := t.TempDir()
	p := filepath.Join(dir, "s")
	w, _ := seg.NewWriter(p, seg.Info{RowWidth: 4, DocOff: 0})
	for i := range uint32(10000) {
		w.Add(nil, binary.BigEndian.AppendUint32(nil, i))
	}
	w.Finish()
	f, _ := seg.Open(p)
	defer f.Close()
	var live Bitmap
	for i := uint32(0); i < 10000; i += 2 {
		live.Put(i, true)
	}
	d, n := probeDead(f, live)
	if n < staleProbe || n >= 10000 || d != 0 {
		t.Fatal(d, n)
	}
	d, n = probeDead(f, Bitmap{})
	if d != n {
		t.Fatal(d, n)
	}
}
