//ff:func feature=index type=reader control=sequence
//ff:what Index.load — read a generation: manifest, sets against the archive, spec, document table, live bitmap, segments, pack paths
package index

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// load fills ix for generation gen.
func (ix *Index) load(a *archive.Archive, gen uint64) error {
	m, err := readManifest(ix.Dir, gen)
	if err != nil {
		return err
	}
	ix.M = m
	if err := checkSets(a, m.Sets); err != nil {
		return err
	}
	if ix.Spec, err = spec.Load(filepath.Join(ix.Dir, specName(m.Spec))); err != nil {
		return err
	}
	if ix.Docs, err = loadDocs(ix.Dir, m); err != nil {
		return err
	}
	if ix.Live, err = readBitmap(filepath.Join(ix.Dir, m.Live)); err != nil {
		return err
	}
	if ix.paths, err = packPaths(a, m.Packs); err != nil {
		return err
	}
	ix.packs = make([]*os.File, len(ix.paths))
	ix.Segs, err = openSegs(ix.Dir, m)
	return err
}
