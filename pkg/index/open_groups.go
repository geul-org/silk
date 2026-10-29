//ff:func feature=index type=reader control=sequence
//ff:what openGroups — open a merge's input segments and list their groups in key order; the files are returned even on a grouping error so the caller closes them
package index

import "github.com/geul-org/silk/pkg/seg"

// openGroups is mergeTask's input.
func openGroups(dir string, rs []SegRef) ([]*seg.File, []seg.Group, error) {
	fs, err := openRefs(dir, rs)
	if err != nil {
		return nil, nil, err
	}
	gs, err := seg.GroupsOf(fs)
	return fs, gs, err
}
