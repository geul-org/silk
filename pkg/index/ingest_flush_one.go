//ff:func feature=index type=writer control=sequence
//ff:what ingest.flushOne — one bucket as a new segment (sorted run or appended segment), recorded for its tag; nothing for an empty bucket
package index

import "path/filepath"

// flushOne writes bucket i.
func (g *ingest) flushOne(i int) error {
	b := g.buckets[i]
	if b.Rows == 0 {
		return nil
	}
	name := g.m.newName(dirSeg, ".sseg")
	info, err := b.Write(filepath.Join(g.dir, name), g.sorted[i])
	if err != nil {
		return err
	}
	b.Reset()
	g.written += info.Bytes
	g.out[i] = append(g.out[i], SegRef{File: name, Rows: info.Rows + info.Other, Bytes: info.Bytes})
	return nil
}
