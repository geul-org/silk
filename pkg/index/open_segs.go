//ff:func feature=index type=reader control=iteration dimension=1
//ff:what openSegs — open segments
package index

import "github.com/geul-org/silk/pkg/seg"

func openSegs(paths []string) ([]*seg.File, error) {
	out := make([]*seg.File, 0, len(paths))
	for _, p := range paths {
		f, err := seg.Open(p)
		if err != nil {
			closeSegs(out)
			return nil, err
		}
		out = append(out, f)
	}
	return out, nil
}
