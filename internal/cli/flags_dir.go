//ff:func feature=cli type=helper control=sequence
//ff:what flags.dir — the index directory of an archive
package cli

import "path/filepath"

func (f *flags) dir(archive string) string {
	if f.index != "" {
		return f.index
	}
	return filepath.Join(archive, ".silk")
}
