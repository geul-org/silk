//ff:func feature=query type=builder control=iteration dimension=1
//ff:what planCond — map a condition onto its section: fields must be fixed positions of the layout; the field at the partition key path gives the partition slot, the one at the sort key path the sort slot
package query

import (
	"fmt"
	"slices"

	"github.com/geul-org/silk/pkg/index"
	"github.com/geul-org/silk/pkg/spec"
)

// planCond fails for a kind the index does not cover or a field path the
// layout has no fixed position for.
func planCond(ix *index.Index, c Cond) (CondPlan, error) {
	sec, t, ok := ix.Section(spec.TagWord(c.Kind, c.Codec))
	if !ok {
		return CondPlan{}, fmt.Errorf("%w: %s", ErrNoSection, c.Name)
	}
	l, _ := spec.LayoutOf(c.Kind, c.Codec)
	p := CondPlan{Cond: c, Section: sec}
	for _, f := range c.Fields {
		if err := spec.CheckKey(l, spec.Key{Path: f.Path, Width: 2}, false); err != nil {
			return CondPlan{}, fmt.Errorf("%w: %s: %v", ErrSyntax, c.Text, err)
		}
		if len(t.Partition.Path) > 0 && slices.Equal(f.Path, t.Partition.Path) {
			p.Sel.Part, p.ByPart = spec.AppendSlot(nil, f.Value, t.Partition.Width), true
		}
		if len(t.Sort.Path) > 0 && slices.Equal(f.Path, t.Sort.Path) {
			p.Sel.Sort, p.BySort = spec.AppendSlot(nil, f.Value, t.Sort.Width), true
		}
	}
	return p, nil
}
