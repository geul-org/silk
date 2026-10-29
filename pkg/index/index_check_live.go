//ff:func feature=index type=rule control=iteration dimension=1
//ff:what Index.checkLive — the live document table equals the archive's live versions (identity and place); each archive document's sequence number
package index

import "fmt"

// checkLive returns seqs[i] for docs[i].
func (ix *Index) checkLive(docs []liveDoc) ([]uint32, error) {
	d := diffDocs(ix.M, ix.Docs, ix.Live, docs)
	if len(d.added) > 0 || len(d.killed) > 0 || ix.Live.Count() != len(docs) {
		return nil, fmt.Errorf("%w: %d archive versions not indexed, %d indexed versions not in the archive", ErrMismatch, len(d.added), len(d.killed))
	}
	seqs := make([]uint32, len(docs))
	for i, k := range d.kept {
		seqs[i] = k.seq
	}
	return seqs, nil
}
