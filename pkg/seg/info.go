//ff:type feature=seg type=model
//ff:what Info — a segment's footer: TAG kind, key widths, sorted flag, tag spec hash, counts, body hash, file size
package seg

// Info describes a segment. Rows counts the rows of the partitions in the
// table, Other the rows of the other partition. Bytes is the file size (not
// stored).
type Info struct {
	Tag       uint16
	PartWidth int
	SortWidth int
	Sorted    bool
	Spec      [32]byte
	Parts     uint32
	Rows      uint64
	Other     uint64
	Sum       [32]byte
	Bytes     int64
}
