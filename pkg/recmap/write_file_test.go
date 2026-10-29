//ff:func feature=recmap type=writer control=sequence
//ff:what TestWriteFile — the file opens back; encode, directory, temporary-file and rename failures surface
package recmap

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteFile(t *testing.T) {
	m, _, _ := sampleMap(t)
	dir := t.TempDir()
	p := filepath.Join(dir, "a", "m.srm")
	if err := WriteFile(p, m); err != nil {
		t.Fatal(err)
	}
	if f, err := Open(p); err != nil || f.Header != m.Header {
		t.Fatal(err)
	} else {
		f.Close()
	}
	bad := *m
	bad.Header.Issuer = "x"
	if err := WriteFile(filepath.Join(dir, "b.srm"), &bad); err == nil {
		t.Fatal("encode error must fail")
	}
	os.WriteFile(filepath.Join(dir, "file"), nil, 0o644)
	if err := WriteFile(filepath.Join(dir, "file", "m.srm"), m); err == nil {
		t.Fatal("parent is a file")
	}
	os.MkdirAll(filepath.Join(dir, "c.srm.tmp"), 0o755)
	if err := WriteFile(filepath.Join(dir, "c.srm"), m); err == nil {
		t.Fatal("temporary path is a directory")
	}
	os.MkdirAll(filepath.Join(dir, "d.srm", "x"), 0o755)
	if err := WriteFile(filepath.Join(dir, "d.srm"), m); err == nil {
		t.Fatal("target is a non-empty directory")
	}
}
