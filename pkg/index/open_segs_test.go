//ff:func feature=index type=reader control=sequence
//ff:what TestOpenSegs — one file list per section; a missing segment fails
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestOpenSegs(t *testing.T) {
	_, _, dir := built(t)
	m, _ := CurrentManifest(dir)
	ss, err := openSegs(dir, m)
	if err != nil || len(ss) != 7 || len(ss[0]) != 1 {
		t.Fatal(err)
	}
	closeSegs(ss)
	os.Remove(filepath.Join(dir, m.Sections[6].Runs[0].File))
	if _, err := openSegs(dir, m); err == nil {
		t.Fatal("missing segment")
	}
}
