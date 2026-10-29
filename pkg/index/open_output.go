//ff:func feature=index type=writer control=sequence
//ff:what openOutput — the Writer of a merge's sorted run (new, or resumed from a checkpoint), the key to resume after and the bytes already written
package index

import (
	"path/filepath"

	"github.com/geul-org/silk/pkg/seg"
)

// openOutput takes the TAG, widths and spec hash from an input's footer.
func openOutput(dir string, t Task, in seg.Info, cp *checkpoint) (*seg.Writer, []byte, int64, error) {
	if cp != nil {
		w, err := seg.ResumeWriter(cp.State)
		return w, cp.Last, cp.State.Written, err
	}
	info := seg.Info{Tag: in.Tag, PartWidth: in.PartWidth, SortWidth: in.SortWidth, Sorted: true, Spec: in.Spec}
	w, err := seg.NewWriter(filepath.Join(dir, t.Output), info, t.Min)
	return w, nil, 0, err
}
