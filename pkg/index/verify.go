//ff:func feature=index type=rule control=iteration dimension=1
//ff:what Verify — check an index against its archive: sets, every segment (hash, order, footer against the spec), the live documents, each section's live rows against rows rebuilt from the documents, and dead counts
package index

import (
	"context"

	"github.com/geul-org/silk/pkg/archive"
)

// Verify reads every live document again; it is as expensive as a build.
func Verify(ctx context.Context, a *archive.Archive, dir string) (VerifyReport, error) {
	ix, err := Open(a, dir)
	if err != nil {
		return VerifyReport{}, err
	}
	defer ix.Close()
	docs, err := liveDocs(ctx, a)
	if err != nil {
		return VerifyReport{}, err
	}
	seqs, err := ix.checkLive(docs)
	if err != nil {
		return VerifyReport{}, err
	}
	want, err := expectRows(ix, docs, seqs)
	if err != nil {
		return VerifyReport{}, err
	}
	rep := VerifyReport{Generation: ix.M.Generation, Sections: len(ix.M.Sections), Docs: len(docs)}
	for i := range ix.M.Sections {
		n, err := ix.checkSection(i, want[i])
		if err != nil {
			return rep, err
		}
		rep.Segments += len(ix.Segs[i])
		rep.Rows += n
	}
	return rep, nil
}
