//ff:func feature=index type=rule control=sequence
//ff:what TestIndexCheckLive — the archive's documents map to their sequence numbers; a document missing from the index fails
package index

import (
	"context"
	"errors"
	"testing"
)

func TestIndexCheckLive(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	docs, _ := liveDocs(context.Background(), a)
	seqs, err := ix.checkLive(docs)
	if err != nil || len(seqs) != 4 || ix.Docs[seqs[2]].ID != docs[2].Doc.ID {
		t.Fatal(err, seqs)
	}
	ix.Live.Put(3, false)
	if _, err := ix.checkLive(docs); !errors.Is(err, ErrMismatch) {
		t.Fatal(err)
	}
}
