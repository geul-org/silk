//ff:func feature=archive type=reader control=sequence
//ff:what Set.Packs — the set's data packs holding live version documents, each with its documents in offset order
package archive

import (
	"context"
	"os"
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/pack"
)

// Packs reads the manifest's OBJECTS (for object names), opens the generation
// pinned to s.Manifest and replays its inventories (Reader.State) to place
// every live version document. A compressed one is ErrCompressed.
func (s Set) Packs(ctx context.Context) ([]Pack, error) {
	b, err := os.ReadFile(filepath.Join(s.Dir, "m", s.Manifest+".gpm"))
	if err != nil {
		return nil, err
	}
	objs, err := parseObjects(b)
	if err != nil {
		return nil, err
	}
	r, err := pack.Open(ctx, pack.FileStore{Root: s.Dir}, pack.Options{Manifest: s.Manifest})
	if err != nil {
		return nil, err
	}
	st, err := r.State(ctx)
	if err != nil {
		return nil, err
	}
	by, err := groupDocs(st, len(objs))
	if err != nil {
		return nil, err
	}
	return makePacks(s, objs, by)
}
