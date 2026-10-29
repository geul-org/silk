//ff:func feature=index type=writer control=sequence
//ff:what writeGeneration — store the spec, a document table file for new entries and the live bitmap, record the archive sets, commit and collect
package index

import (
	"os"
	"path/filepath"

	"github.com/park-jun-woo/silk/pkg/archive"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// writeGeneration ends a build or update; rep gets the bytes written.
func writeGeneration(dir string, m *Manifest, sp spec.Spec, a *archive.Archive, entries []DocEntry, live Bitmap, rep *Report) error {
	m.Spec, m.Sets = sp.Hash(), setRefs(a)
	if _, err := os.Stat(filepath.Join(dir, specName(m.Spec))); err != nil {
		n, err := writeAtomic(filepath.Join(dir, specName(m.Spec)), sp.Encode())
		rep.Written += n
		if err != nil {
			return err
		}
	}
	if len(entries) > 0 {
		first := uint32(0)
		for _, d := range m.Docs {
			first += d.Count
		}
		f := DocFile{File: m.newName(dirDoc, ".sdoc"), First: first, Count: uint32(len(entries))}
		n, err := writeAtomic(filepath.Join(dir, f.File), encodeDocs(first, entries))
		rep.Written += n
		if err != nil {
			return err
		}
		m.Docs = append(m.Docs, f)
	}
	m.Live = m.newName(dirLive, ".sbit")
	n, err := writeAtomic(filepath.Join(dir, m.Live), encodeBitmap(live))
	rep.Written += n
	if err != nil {
		return err
	}
	n, err = commit(dir, m)
	rep.Written += n
	rep.Generation = m.Generation
	collect(dir, *m)
	return err
}
