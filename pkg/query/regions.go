//ff:func feature=query type=builder control=iteration dimension=1
//ff:what regions — the partitions a plan reads of one segment
package query

import "github.com/geul-org/silk/pkg/seg"

func regions(p *Plan, f *seg.File) []seg.Region {
	if len(p.Parts) > 0 {
		return fixedRegions(p, f)
	}
	var out []seg.Region
	for i := uint64(0); i < f.Parts; i++ {
		_, r := f.PartAt(i)
		out = append(out, r)
	}
	return out
}
