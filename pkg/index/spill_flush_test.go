//ff:func feature=index type=scanner control=sequence
//ff:what TestSpillFlush — a run per flush; a row that does not fit the shape and a failing finish stop
package index

import (
	"fmt"
	"os"
	"path/filepath"
	"sync/atomic"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestSpillFlush(t *testing.T) {
	var n atomic.Uint64
	dir := t.TempDir()
	info := seg.Info{RowWidth: 8, DocOff: 4}
	s := newSpill(info, dir, &n)
	s.buf.Add(nil, make([]byte, 8))
	if err := s.flush(); err != nil || len(s.runs) != 1 || s.buf.Rows() != 0 {
		t.Fatal(err)
	}
	wrong := newSpill(seg.Info{RowWidth: 4, DocOff: 0}, dir, &n)
	wrong.buf = seg.NewBuffer(info)
	wrong.buf.Add(nil, make([]byte, 8))
	if err := wrong.flush(); err == nil {
		t.Fatal("shape")
	}
	blocked := newSpill(info, dir, &n)
	blocked.buf.Add(nil, make([]byte, 8))
	name := filepath.Join(dir, fmt.Sprintf("%d.run", n.Load()+1))
	os.MkdirAll(filepath.Join(name, "x"), 0o755)
	if err := blocked.flush(); err == nil {
		t.Fatal("rename over a directory")
	}
}
