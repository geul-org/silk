//ff:func feature=seg type=encoder control=sequence
//ff:what encodeFooter — the 128-byte footer: magic, version, flags, TAG, key widths, counts, tag spec hash, body hash, magic
package seg

import "encoding/binary"

// encodeFooter lays out docs/segments.md "Footer".
func encodeFooter(i Info) []byte {
	be := binary.BigEndian
	b := append(make([]byte, 0, footerSize), magic...)
	flags := uint16(0)
	if i.Sorted {
		flags = flagSorted
	}
	b = be.AppendUint16(be.AppendUint16(be.AppendUint16(b, version), flags), i.Tag)
	b = be.AppendUint16(be.AppendUint16(be.AppendUint16(b, uint16(i.PartWidth)), uint16(i.SortWidth)), 0)
	b = be.AppendUint64(be.AppendUint64(be.AppendUint32(b, i.Parts), i.Rows), i.Other)
	b = append(append(b, i.Spec[:]...), i.Sum[:]...)
	b = append(b, make([]byte, footerSize-8-len(b))...)
	return append(b, magic...)
}
