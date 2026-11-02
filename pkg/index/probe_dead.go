//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what probeDead — dead and total rows among evenly spaced probes
package index

import "github.com/geul-org/silk/pkg/seg"

func probeDead(f *seg.File, live Bitmap) (int, int) {
	r := f.All()
	step := max(r.Len()/staleProbe, 1)
	dead, n := 0, 0
	for i := 0; i < r.Len(); i += step {
		n++
		if !live.Get(f.Doc(r.Row(i))) {
			dead++
		}
	}
	return dead, n
}
