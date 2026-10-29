//ff:func feature=index type=helper control=sequence
//ff:what TestManifestNewName — each call reserves the next file number
package index

import "testing"

func TestManifestNewName(t *testing.T) {
	m := Manifest{NextFile: 9}
	a, b := m.newName(dirSeg, ".sseg"), m.newName(dirDoc, ".sdoc")
	if a != "seg/000000000000000a.sseg" || b != "doc/000000000000000b.sdoc" || m.NextFile != 11 {
		t.Fatal(a, b)
	}
}
