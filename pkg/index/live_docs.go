//ff:func feature=index type=reader control=iteration dimension=1
//ff:what liveDocs — every live version document of an archive: set by set, data pack by pack, by offset
package index

import (
	"context"

	"github.com/park-jun-woo/silk/pkg/archive"
)

// liveDocs lists what the archive's current generations hold (archive.Set.Packs).
func liveDocs(ctx context.Context, a *archive.Archive) ([]liveDoc, error) {
	var out []liveDoc
	for _, s := range a.Sets {
		packs, err := s.Packs(ctx)
		if err != nil {
			return nil, err
		}
		out = appendPackDocs(out, s.Name, packs)
	}
	return out, nil
}
