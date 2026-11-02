//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what Index.printRows — one partition's rows into a fingerprint
package index

import "github.com/geul-org/silk/pkg/seg"

func (ix *Index) printRows(p *fingerprint, f *seg.File, key []byte, r seg.Region) {
	for i := 0; i < r.Len(); i++ {
		row := r.Row(i)
		seq := f.Doc(row)
		if ix.Live.Get(seq) {
			p.add(key, row[:f.DocOff], []byte(ix.Path(seq)), row[f.DocOff+4:])
		}
	}
}
