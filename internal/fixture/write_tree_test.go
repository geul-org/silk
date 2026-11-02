//ff:func feature=fixture type=scanner control=sequence
//ff:what TestWriteTree — unwritable places and an unencodable document fail the test
package fixture

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/pkg/archive"
)

func TestWriteTree(t *testing.T) {
	file := filepath.Join(t.TempDir(), "file")
	os.WriteFile(file, nil, 0o644)
	if !fails(t, func(tb testing.TB) { writeTree(tb, file, []Doc{{ID: 1, Version: 1}}) }) {
		t.Fatal("under a file")
	}
	bad := []core.Record{{Kind: core.KindValue, Codec: 0xFF}}
	if !fails(t, func(tb testing.TB) { writeTree(tb, t.TempDir(), []Doc{{ID: 1, Version: 1, Body: bad}}) }) {
		t.Fatal("bad record")
	}
	dir := t.TempDir()
	p := filepath.Join(dir, ".well-known", "geul", archive.VersionPath(issuer, 1, 1))
	os.MkdirAll(p, 0o755)
	if !fails(t, func(tb testing.TB) { writeTree(tb, dir, []Doc{{ID: 1, Version: 1}}) }) {
		t.Fatal("path is a directory")
	}
	out := t.TempDir()
	os.WriteFile(filepath.Join(out, ".well-known"), nil, 0o644)
	if !fails(t, func(tb testing.TB) { Repack(tb, out, []Doc{{ID: 1, Version: 1}}, nil) }) {
		t.Fatal("output under a file")
	}
}
