//ff:func feature=query type=builder control=iteration dimension=1
//ff:what fixedRegions — the partitions a predicate named, when present
package query

import "github.com/geul-org/silk/pkg/seg"

func fixedRegions(p *Plan, f *seg.File) []seg.Region {
	var out []seg.Region
	for _, k := range p.Parts {
		if r, ok := f.Part(k); ok {
			out = append(out, r)
		}
	}
	return out
}
