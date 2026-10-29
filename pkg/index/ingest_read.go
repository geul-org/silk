//ff:func feature=index type=reader control=sequence
//ff:what ingest.readDoc — a document's bytes from its data pack (pack files kept open)
package index

import "os"

// readDoc reads d into the ingest's buffer.
func (g *ingest) readDoc(d liveDoc) ([]byte, error) {
	f := g.files[d.Path]
	if f == nil {
		var err error
		if f, err = os.Open(d.Path); err != nil {
			return nil, err
		}
		g.files[d.Path] = f
	}
	if uint64(cap(g.buf)) < d.Doc.Len {
		g.buf = make([]byte, d.Doc.Len)
	}
	g.buf = g.buf[:d.Doc.Len]
	if _, err := f.ReadAt(g.buf, int64(d.Doc.Off)); err != nil {
		return nil, err
	}
	g.read += int64(d.Doc.Len)
	return g.buf, nil
}
