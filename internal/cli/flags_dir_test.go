//ff:func feature=cli type=scanner control=sequence
//ff:what TestFlagsDir — the given directory, or .silk in the archive
package cli

import (
	"path/filepath"
	"testing"
)

func TestFlagsDir(t *testing.T) {
	if d := (&flags{index: "x"}).dir("a"); d != "x" {
		t.Fatal(d)
	}
	if d := (&flags{}).dir("a"); d != filepath.Join("a", ".silk") {
		t.Fatal(d)
	}
}
