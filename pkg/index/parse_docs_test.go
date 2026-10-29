//ff:func feature=index type=decoder control=sequence
//ff:what TestParseDocs — entries with pair ranges; a range past the pairs area or nonzero reserved bytes fail
package index

import (
	"testing"
)

func TestParseDocs(t *testing.T) {
	b := encodeDocs(0, []DocEntry{{ID: 4, Counts: []TagCount{{7, 2}}}})[docsHead:]
	e, p := b[:docSize], b[docSize:]
	if ds, err := parseDocs(e, p); err != nil || ds[0].ID != 4 || ds[0].Counts[0].N != 2 {
		t.Fatal(err, ds)
	}
	if _, err := parseDocs(e, nil); err == nil {
		t.Fatal("pairs past the end")
	}
	r := append([]byte{}, e...)
	r[39] = 1
	if _, err := parseDocs(r, p); err == nil {
		t.Fatal("reserved bytes")
	}
}
