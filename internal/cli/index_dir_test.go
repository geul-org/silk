//ff:func feature=cli type=helper control=sequence
//ff:what TestIndexDir — the flag wins; otherwise .silk in the data folder, also when given its .well-known/geul
package cli

import (
	"path/filepath"
	"testing"
)

func TestIndexDir(t *testing.T) {
	if indexDir("/i", "/d") != "/i" {
		t.Fatal("flag")
	}
	if got := indexDir("", "/d"); got != filepath.Join("/d", ".silk") {
		t.Fatal(got)
	}
	if got := indexDir("", "/d/.well-known/geul/"); got != filepath.Join("/d", ".silk") {
		t.Fatal(got)
	}
}
