//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what runner.fileRows — one segment: silk patterns, or partitions and prefixes; collecting a row cannot fail
package query

import "github.com/geul-org/silk/pkg/seg"

func (r *runner) fileRows(p *Plan, f *seg.File) []Hit {
	var hits []Hit
	emit := func(row []byte) error {
		if h, ok := r.hit(p, f, row); ok {
			hits = append(hits, h)
		}
		return nil
	}
	for _, reg := range regions(p, f) {
		_ = scanRegion(p, reg, emit)
	}
	return hits
}
