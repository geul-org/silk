//ff:func feature=index type=rule control=sequence
//ff:what TestSameLoc — same set, pack, offset and length only
package index

import (
	"testing"

	"github.com/geul-org/silk/pkg/archive"
)

func TestSameLoc(t *testing.T) {
	m := Manifest{Packs: []PackRef{{Set: "s", Hash: "h"}}}
	e := DocEntry{Off: 4, Len: 6}
	d := liveDoc{Set: "s", Pack: "h", Doc: archive.Doc{Off: 4, Len: 6}}
	if !sameLoc(m, e, d) {
		t.Fatal("same")
	}
	d.Doc.Len = 8
	if sameLoc(m, e, d) {
		t.Fatal("moved")
	}
}
