//ff:func feature=index type=rule control=sequence
//ff:what checkSeg — a segment's bytes verify and its footer names the tag's TAG WORD, key widths and spec hash, sorted exactly when it is a run
package index

import (
	"fmt"

	"github.com/park-jun-woo/silk/pkg/seg"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// checkSeg compares f's footer with t.
func checkSeg(f *seg.File, t spec.Tag, run bool) error {
	if err := f.Verify(); err != nil {
		return err
	}
	i := f.Info
	if i.Tag != spec.TagWord(t.Kind, t.Codec) || i.PartWidth != t.Partition.Width || i.SortWidth != t.Sort.Width || i.Spec != t.Hash() || i.Sorted != run {
		return fmt.Errorf("%w: %s does not match tag %s", ErrMismatch, f.Path, t.Name)
	}
	return nil
}
