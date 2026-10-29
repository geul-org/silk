//ff:func feature=index type=helper control=sequence
//ff:what TestAppendPackDocs — every document of every pack, with the pack's hash and path
package index

import (
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

func TestAppendPackDocs(t *testing.T) {
	ps := []archive.Pack{{Hash: "a", Path: "pa", Docs: []archive.Doc{{ID: 1}, {ID: 2}}}, {Hash: "b", Docs: []archive.Doc{{ID: 3}}}}
	ds := appendPackDocs(nil, "s", ps)
	if len(ds) != 3 || ds[1].Pack != "a" || ds[1].Path != "pa" || ds[2].Doc.ID != 3 || ds[2].Set != "s" {
		t.Fatal(ds)
	}
}
