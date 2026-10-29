//ff:func feature=index type=helper control=iteration dimension=2
//ff:what appendPackDocs — the documents of a set's data packs as liveDocs, pack by pack
package index

import "github.com/geul-org/silk/pkg/archive"

// appendPackDocs flattens packs.
func appendPackDocs(out []liveDoc, set string, packs []archive.Pack) []liveDoc {
	for _, p := range packs {
		for _, d := range p.Docs {
			out = append(out, liveDoc{Set: set, Pack: p.Hash, Path: p.Path, Doc: d})
		}
	}
	return out
}
