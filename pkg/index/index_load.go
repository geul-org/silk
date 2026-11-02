//ff:func feature=index type=reader control=sequence
//ff:what Index.load — manifest, staleness check, bitmap, table, grammar and sections
package index

import "slices"

func (ix *Index) load(g uint64) error {
	m, err := readManifest(ix.Dir, g)
	if err != nil {
		return err
	}
	ix.Manifest = m
	if !slices.Equal(m.Sets, setRefs(ix.ar)) {
		return ErrStale
	}
	if ix.Live, err = loadLiveOf(ix.Dir, m); err != nil {
		return err
	}
	if ix.Docs, err = openDocs(ix.Dir, &m); err != nil {
		return err
	}
	if ix.Grammar, err = loadGrammar(ix.Dir, m.Grammar); err != nil {
		return err
	}
	return ix.loadSections()
}
