//ff:func feature=index type=scanner control=sequence
//ff:what TestSpillFinish — one run goes through unmerged; several merge
package index

import (
	"path/filepath"
	"sync/atomic"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestSpillFinish(t *testing.T) {
	var n atomic.Uint64
	dir := t.TempDir()
	info := seg.Info{RowWidth: 8, DocOff: 4}
	s := newSpill(info, dir, &n)
	s.buf.Add(nil, make([]byte, 8))
	p, err := s.finish(2)
	if err != nil || p == "" || s.name() == p {
		t.Fatal(p, err)
	}
	b := newSpill(info, filepath.Join(p, "x"), &n)
	b.buf.Add(nil, make([]byte, 8))
	if _, err := b.finish(2); err == nil {
		t.Fatal("under a file")
	}
}
