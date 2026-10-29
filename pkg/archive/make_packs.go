//ff:func feature=archive type=builder control=iteration dimension=1
//ff:what makePacks — one Pack per data object holding documents, in object order, documents sorted by offset
package archive

import (
	"fmt"
	"path/filepath"
	"sort"
)

// makePacks turns the grouped documents into Packs. A document placed in an
// object that is not a data pack is ErrManifest.
func makePacks(s Set, objs []Object, by map[uint32][]Doc) ([]Pack, error) {
	var out []Pack
	for i, o := range objs {
		docs := by[uint32(i)]
		if len(docs) == 0 {
			continue
		}
		if o.Kind != kindData {
			return nil, fmt.Errorf("%w: documents in object %d of kind %d", ErrManifest, i, o.Kind)
		}
		sort.Slice(docs, func(a, b int) bool { return docs[a].Off < docs[b].Off })
		out = append(out, Pack{Set: s, Hash: o.Hash, Path: filepath.Join(s.Dir, "o", o.Hash+".gpk"), Size: o.Size, Docs: docs})
	}
	return out, nil
}
