//ff:func feature=index type=reader control=iteration dimension=1
//ff:what openRefs — open segment files by reference; on error the ones opened are closed
package index

import (
	"path/filepath"

	"github.com/park-jun-woo/silk/pkg/seg"
)

// openRefs opens rs under dir.
func openRefs(dir string, rs []SegRef) ([]*seg.File, error) {
	out := make([]*seg.File, 0, len(rs))
	for _, r := range rs {
		f, err := seg.Open(filepath.Join(dir, r.File))
		if err != nil {
			closeSegs([][]*seg.File{out})
			return nil, err
		}
		out = append(out, f)
	}
	return out, nil
}
