//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestReadStat — counts of a built index; a missing index or files fail
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/pkg/decl"
)

func TestReadStat(t *testing.T) {
	_, dir, _ := built(t, 3)
	st, err := ReadStat(dir)
	if err != nil || st.Docs != 3 || st.Live != 3 || len(st.Sections) != 3 || st.Sections[0].Rows != 3 || st.Sections[0].Kind != decl.KindSilk {
		t.Fatal(st, err)
	}
	if st.MaxDoc == 0 || st.Sections[0].Bytes == 0 {
		t.Fatal(st)
	}
	m := current(t, dir)
	if _, err := ReadStat(t.TempDir()); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	bad := t.TempDir()
	os.WriteFile(filepath.Join(bad, currentFile), []byte("9"), 0o644)
	if _, err := ReadStat(bad); err == nil {
		t.Fatal("no manifest")
	}
	if b, err := loadLiveOf(dir, Manifest{}); err != nil || b.N != 0 {
		t.Fatal(b, err)
	}
	if _, err := statSection(dir, Section{Decl: "none"}, Bitmap{}); err == nil {
		t.Fatal("no declaration")
	}
	if _, err := staleShare(dir, Section{Runs: []SegRef{{File: "none"}}}, Bitmap{}); err == nil {
		t.Fatal("no segment")
	}
	for _, r := range m.Sections[0].Runs {
		os.Remove(filepath.Join(dir, r.File))
	}
	if _, err := ReadStat(dir); err == nil {
		t.Fatal("no segment")
	}
	os.Remove(filepath.Join(dir, m.Live))
	if _, err := ReadStat(dir); err == nil {
		t.Fatal("no bitmap")
	}
}
