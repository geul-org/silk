//ff:func feature=index type=writer control=iteration dimension=1
//ff:what TestWriteGeneration — spec stored once, a document table continuing the sequence, bitmap, sets and a commit; blocked spec, document, bitmap or manifest paths fail
package index

import (
	"testing"
)

func TestWriteGeneration(t *testing.T) {
	a, _, _ := built(t)
	d := t.TempDir()
	m := Manifest{Docs: []DocFile{{Count: 3}}}
	var rep Report
	if err := writeGeneration(d, &m, smallSpec(), a, []DocEntry{{ID: 1}}, allLive(4), &rep); err != nil {
		t.Fatal(err)
	}
	if m.Docs[1].First != 3 || m.Generation != 1 || rep.Generation != 1 || len(m.Sets) != len(a.Sets) || m.Spec != smallSpec().Hash() {
		t.Fatal(m)
	}
	w := rep.Written
	if err := writeGeneration(d, &m, smallSpec(), a, nil, allLive(4), &rep); err != nil || len(m.Docs) != 2 || rep.Written-w > w {
		t.Fatal(err)
	}
	for _, sub := range []string{dirManifest, dirLive, dirDoc, dirSpec} {
		block(d, sub)
		if err := writeGeneration(d, &m, smallSpec(), a, []DocEntry{{}}, allLive(1), &rep); err == nil {
			t.Fatal(sub)
		}
	}
}
