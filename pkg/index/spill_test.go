//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestSpill — runs spill, merge down level by level, and an empty spill has no path
package index

import (
	"errors"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestSpill(t *testing.T) {
	dir := t.TempDir()
	var n atomic.Uint64
	info := seg.Info{RowWidth: 8, DocOff: 4}
	s := newSpill(info, dir, &n)
	if p, err := s.finish(2); p != "" || err != nil {
		t.Fatal(p, err)
	}
	s = newSpill(info, dir, &n)
	for i := range 5 {
		s.buf = seg.NewBuffer(info)
		s.buf.Add(nil, []byte{0, 0, 0, byte(i), 0, 0, 0, 1})
		if err := s.flush(); err != nil {
			t.Fatal(err)
		}
	}
	s.buf = seg.NewBuffer(info)
	p, err := s.finish(2)
	if err != nil {
		t.Fatal(err)
	}
	f, err := seg.Open(p)
	if err != nil || f.Rows != 5 {
		t.Fatal(err)
	}
	f.Close()
	if es, _ := os.ReadDir(dir); len(es) != 1 {
		t.Fatal(es)
	}
	if err := (&spill{buf: seg.NewBuffer(info)}).flush(); err != nil {
		t.Fatal(err)
	}
	bad := newSpill(info, filepath.Join(p, "x"), &n)
	bad.buf.Add(nil, make([]byte, 8))
	if err := bad.flush(); err == nil {
		t.Fatal("under a file")
	}
	if _, err := bad.finish(2); err == nil {
		t.Fatal("finish under a file")
	}
	wrong := newSpill(seg.Info{RowWidth: 8, DocOff: 4, PartWidth: 2}, dir, &n)
	wrong.buf = seg.NewBuffer(info)
	wrong.buf.Add(nil, make([]byte, 8))
	if err := wrong.flush(); !errors.Is(err, seg.ErrFormat) && err == nil {
		t.Fatal(err)
	}
	if fanIn(1) != minFanIn || fanIn(1<<40) != maxFanIn || fanIn(8<<20) != 8 {
		t.Fatal(fanIn(8 << 20))
	}
	if threadsOf(Options{}) != 1 {
		t.Fatal("threads")
	}
}
