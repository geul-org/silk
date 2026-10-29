//ff:func feature=index type=writer control=sequence
//ff:what TestWriteCheckpoints — writes compact.json; an empty list removes it
package index

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteCheckpoints(t *testing.T) {
	d := t.TempDir()
	if n, err := writeCheckpoints(d, []checkpoint{{Section: "S"}}); err != nil || n == 0 {
		t.Fatal(err)
	}
	if n, err := writeCheckpoints(d, nil); err != nil || n != 0 {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(d, checkpointFile)); !os.IsNotExist(err) {
		t.Fatal(err)
	}
}
