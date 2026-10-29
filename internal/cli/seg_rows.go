//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what segRows — rows of a section's runs and appended segments
package cli

import "github.com/geul-org/silk/pkg/index"

// segRows sums SegRef.Rows.
func segRows(s index.Section) uint64 {
	n := uint64(0)
	for _, r := range append(append([]index.SegRef{}, s.Runs...), s.Appends...) {
		n += r.Rows
	}
	return n
}
