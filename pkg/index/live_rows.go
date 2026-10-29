//ff:func feature=index type=reader control=sequence
//ff:what liveRows — a segment's rows of live versions as full rows (partition slot first) appended to got, and its rows of dead versions counted
package index

import "github.com/park-jun-woo/silk/pkg/seg"

// liveRows reads f once.
func liveRows(f *seg.File, live Bitmap, got []byte) ([]byte, uint64, error) {
	dead := uint64(0)
	err := f.Each(func(part, row []byte) error {
		if doc, _ := seg.SeqOf(row); !live.Get(doc) {
			dead++
			return nil
		}
		got = append(append(got, part...), row...)
		return nil
	})
	return got, dead, err
}
