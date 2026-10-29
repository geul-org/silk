//ff:func feature=index type=reader control=sequence
//ff:what TestReadCheckpoints — saved checkpoints read back; a missing or unreadable file gives none
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestReadCheckpoints(t *testing.T) {
	d := t.TempDir()
	if readCheckpoints(d) != nil {
		t.Fatal("none expected")
	}
	writeCheckpoints(d, []checkpoint{{Section: "TRIPLE", Last: []byte{1}}})
	if cs := readCheckpoints(d); len(cs) != 1 || cs[0].Section != "TRIPLE" || cs[0].Last[0] != 1 {
		t.Fatal(cs)
	}
	os.WriteFile(filepath.Join(d, checkpointFile), []byte("{"), 0o644)
	if readCheckpoints(d) != nil {
		t.Fatal("bad JSON")
	}
}
