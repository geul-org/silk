//ff:func feature=index type=encoder control=sequence
//ff:what TestEncodeDocs — entries and their record counts round-trip through readDocs
package index

import (
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func TestEncodeDocs(t *testing.T) {
	in := []DocEntry{
		{Pack: 1, Issuer: 2, Off: 1 << 33, Len: 7, ID: 9, Version: 3, Counts: []TagCount{{0x100A, 4}, {0x1004, 1}}},
		{Pack: 0, Issuer: 0, Off: 5, Len: 6, ID: 1, Version: 1, Counts: []TagCount{}},
	}
	p := filepath.Join(t.TempDir(), "d")
	os.WriteFile(p, encodeDocs(5, in), 0o644)
	first, out, err := readDocs(p)
	if err != nil || first != 5 || !reflect.DeepEqual(out, in) {
		t.Fatal(err, first, out)
	}
}
