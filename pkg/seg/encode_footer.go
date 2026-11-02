//ff:func feature=seg type=encoder control=sequence
//ff:what encodeFooter — the 128-byte footer: magic, version, widths, document offset, counts, hashes, magic
package seg

import "encoding/binary"

func encodeFooter(i Info, parts, rows uint64, sum []byte) []byte {
	b := append(make([]byte, 0, footerSize), magic...)
	b = binary.BigEndian.AppendUint16(b, version)
	b = binary.BigEndian.AppendUint16(b, uint16(i.PartWidth))
	b = binary.BigEndian.AppendUint16(b, uint16(i.RowWidth))
	b = binary.BigEndian.AppendUint16(b, uint16(i.DocOff))
	b = binary.BigEndian.AppendUint64(b, parts)
	b = binary.BigEndian.AppendUint64(b, rows)
	b = append(b, i.Hash[:]...)
	b = append(b, sum...)
	b = append(b, make([]byte, footerSize-len(magic)-len(b))...)
	return append(b, magic...)
}
