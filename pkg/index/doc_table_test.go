//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what TestDocTable — entries round trip across files; a bad path fails; damage shows in check
package index

import (
	"errors"
	"os"
	"path/filepath"
	"testing"
)

func TestDocTable(t *testing.T) {
	dir := t.TempDir()
	e := DocEntry{Pack: 1, Issuer: 2, Off: 3, Len: 4, ID: 5, Version: 6, Enc: 1}
	if parseDocEntry(e.append(nil)) != e {
		t.Fatal("round trip")
	}
	var m Manifest
	for i, name := range []string{"doc/a", "doc/b"} {
		w, err := newDocWriter(filepath.Join(dir, name), uint32(i*2))
		if err != nil {
			t.Fatal(err)
		}
		w.add(DocEntry{ID: uint64(i)})
		w.add(DocEntry{ID: uint64(i + 10)})
		if err := w.finish(); err != nil {
			t.Fatal(err)
		}
		m.Docs = append(m.Docs, DocFile{File: name, First: uint32(i * 2), Count: 2})
	}
	tb, err := openDocs(dir, &m)
	if err != nil {
		t.Fatal(err)
	}
	defer tb.Close()
	if tb.At(3).ID != 11 || tb.At(0).ID != 0 {
		t.Fatal(tb.At(3))
	}
	if err := tb.check(dir, &m); err != nil {
		t.Fatal(err)
	}
	tb.files[1].first = 9
	if err := tb.check(dir, &m); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	m.Docs = append(m.Docs, DocFile{File: "doc/none"})
	if _, err := openDocs(dir, &m); err == nil {
		t.Fatal("missing")
	}
	file := filepath.Join(dir, "f")
	os.WriteFile(file, nil, 0o644)
	if _, err := newDocWriter(filepath.Join(file, "x"), 0); err == nil {
		t.Fatal("under a file")
	}
	os.Mkdir(filepath.Join(dir, "g.tmp"), 0o755)
	if _, err := newDocWriter(filepath.Join(dir, "g"), 0); err == nil {
		t.Fatal("tmp is a directory")
	}
	w, _ := newDocWriter(filepath.Join(dir, "h"), 0)
	w.f.Close()
	if err := w.finish(); err == nil {
		t.Fatal("closed file")
	}
}
