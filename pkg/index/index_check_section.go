//ff:func feature=index type=rule control=iteration dimension=1
//ff:what Index.checkSection — every segment of a section verified and matching the spec tag; its live rows equal the expected rows and its dead rows the dead count
package index

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/park-jun-woo/silk/pkg/seg"
)

// checkSection returns the section's live rows.
func (ix *Index) checkSection(i int, want []byte) (uint64, error) {
	s := ix.M.Sections[i]
	t, ok := ix.Spec.Find(s.Tag)
	if !ok {
		return 0, fmt.Errorf("%w: section %s is not in the spec", ErrMismatch, s.Name)
	}
	var got []byte
	var errs []error
	dead := uint64(0)
	for j, f := range ix.Segs[i] {
		var d uint64
		var err error
		got, d, err = liveRows(f, ix.Live, got)
		errs = append(errs, checkSeg(f, t, j < len(s.Runs)), err)
		dead += d
	}
	if err := errors.Join(errs...); err != nil {
		return 0, err
	}
	w := t.Partition.Width + t.Sort.Width + 8
	if dead != s.Dead || !bytes.Equal(seg.SortRows(got, w), want) {
		return 0, fmt.Errorf("%w: section %s: %d live rows (want %d), %d dead (manifest %d)", ErrMismatch, s.Name, len(got)/w, len(want)/w, dead, s.Dead)
	}
	return uint64(len(got) / w), nil
}
