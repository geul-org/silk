//ff:func feature=index type=reader control=iteration dimension=1
//ff:what openSegs — the segment files of every section (runs, then appended segments)
package index

import "github.com/geul-org/silk/pkg/seg"

// openSegs opens each section; on error the files opened so far are closed.
func openSegs(dir string, m Manifest) ([][]*seg.File, error) {
	out := make([][]*seg.File, 0, len(m.Sections))
	for _, s := range m.Sections {
		fs, err := openRefs(dir, append(append([]SegRef{}, s.Runs...), s.Appends...))
		if err != nil {
			closeSegs(out)
			return nil, err
		}
		out = append(out, fs)
	}
	return out, nil
}
