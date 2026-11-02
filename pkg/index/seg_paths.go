//ff:func feature=index type=helper control=iteration dimension=1
//ff:what segPaths — the full paths of a section's segments
package index

import (
	"path/filepath"
	"slices"
)

func segPaths(dir string, s Section) []string {
	var out []string
	for _, r := range append(slices.Clone(s.Runs), s.Appends...) {
		out = append(out, filepath.Join(dir, r.File))
	}
	return out
}
