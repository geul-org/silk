//ff:func feature=cli type=helper control=sequence
//ff:what indexDir — the -index value, or `.silk` in the archive's data folder (the parent of .well-known when given that directory)
package cli

import "path/filepath"

// indexDir picks the index directory for archive.
func indexDir(flagDir, archive string) string {
	if flagDir != "" {
		return flagDir
	}
	a := filepath.Clean(archive)
	if filepath.Base(a) == "geul" && filepath.Base(filepath.Dir(a)) == ".well-known" {
		a = filepath.Dir(filepath.Dir(a))
	}
	return filepath.Join(a, ".silk")
}
