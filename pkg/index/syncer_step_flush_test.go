//ff:func feature=index type=scanner control=sequence
//ff:what TestSyncerStepFlush — a flush that fails stops the step
package index

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

func TestSyncerStepFlush(t *testing.T) {
	u := &updater{o: Options{Limit: 1}, tmp: t.TempDir()}
	s := &syncer{u: u, iss: map[string]uint32{}}
	s.dw, _ = newDocWriter(filepath.Join(u.tmp, "d"), 0)
	file := filepath.Join(u.tmp, "f")
	os.WriteFile(file, nil, 0o644)
	s.sp = newSpill(pathInfo(), file, &u.n)
	if err := s.step(archive.Step{Issuer: "a"}); err == nil {
		t.Fatal("flush under a file")
	}
	u.o.Limit = 1 << 20
	s.sp = newSpill(pathInfo(), u.tmp, &u.n)
	if err := s.step(archive.Step{Issuer: "a"}); err != nil || s.sp.buf.Rows() != 1 {
		t.Fatal(err)
	}
}
