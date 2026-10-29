//ff:func feature=index type=builder control=sequence
//ff:what TestStartManifest — empty in a new directory; continues generation and file numbers of an existing index
package index

import "testing"

func TestStartManifest(t *testing.T) {
	d := t.TempDir()
	if m := startManifest(d); m.Generation != 0 || m.NextFile != 0 || m.Version != manifestVersion {
		t.Fatal(m)
	}
	cur := Manifest{NextFile: 9, Spec: "x"}
	commit(d, &cur)
	if m := startManifest(d); m.Generation != 1 || m.NextFile != 9 || m.Spec != "" {
		t.Fatal(m)
	}
}
