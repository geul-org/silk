//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what segBytes — bytes of a section's runs and appended segments
package cli

import "github.com/park-jun-woo/silk/pkg/index"

// segBytes sums SegRef.Bytes.
func segBytes(s index.Section) int64 {
	n := int64(0)
	for _, r := range append(append([]index.SegRef{}, s.Runs...), s.Appends...) {
		n += r.Bytes
	}
	return n
}
