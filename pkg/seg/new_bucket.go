//ff:func feature=seg type=builder control=sequence
//ff:what NewBucket — an empty Bucket for a TAG kind's key widths, spec hash and minimum partition
package seg

// NewBucket returns a Bucket whose segments carry info (counts ignored).
func NewBucket(info Info, min int) *Bucket {
	return &Bucket{Info: info, Min: min, parts: map[string][]byte{}}
}
