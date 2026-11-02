//ff:func feature=index type=reader control=sequence
//ff:what infoOf — a segment's shape
package index

import "github.com/geul-org/silk/pkg/seg"

func infoOf(path string) (seg.Info, error) {
	f, err := seg.Open(path)
	if err != nil {
		return seg.Info{}, err
	}
	defer f.Close()
	return f.Info, nil
}
