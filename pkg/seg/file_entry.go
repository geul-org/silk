//ff:func feature=seg type=helper control=sequence
//ff:what File.entry — partition i: key, first row, row count
package seg

import "encoding/binary"

func (f *File) entry(i uint64) ([]byte, uint64, uint64) {
	w := uint64(f.PartWidth) + 16
	e := f.table[i*w : (i+1)*w]
	k := e[:f.PartWidth]
	return k, binary.BigEndian.Uint64(e[f.PartWidth:]), binary.BigEndian.Uint64(e[f.PartWidth+8:])
}
