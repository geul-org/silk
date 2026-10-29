//ff:func feature=index type=reader control=sequence
//ff:what TestIngestReadDoc — a document's bytes (pack kept open, bytes counted); a missing pack or a range past its end fails
package index

import "testing"

func TestIngestReadDoc(t *testing.T) {
	g, _, _, docs := testIngest(t, 1<<20)
	b, err := g.readDoc(docs[0])
	if err != nil || uint64(len(b)) != docs[0].Doc.Len || g.read != int64(len(b)) || len(g.files) != 1 {
		t.Fatal(err)
	}
	g.readDoc(docs[1])
	if len(g.files) != 1 {
		t.Fatal("pack reopened")
	}
	bad := docs[0]
	bad.Doc.Off = 1 << 40
	if _, err := g.readDoc(bad); err == nil {
		t.Fatal("past the end")
	}
	bad.Path = "/nonexistent"
	if _, err := g.readDoc(bad); err == nil {
		t.Fatal("missing pack")
	}
}
