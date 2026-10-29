//ff:type feature=seg type=model
//ff:what region — a run of fixed-width rows in a segment: byte offset of the first row, row count, row width
package seg

// region is either one partition or the other partition.
type region struct {
	off int64
	n   uint64
	w   int
}
