//ff:func feature=index type=helper control=sequence
//ff:what openBroken — test helper: build an index, damage it with fn (given the directory and current manifest), and return Open's error
package index

import "testing"

// openBroken isolates one failure of Index.load.
func openBroken(t *testing.T, fn func(dir string, m Manifest)) error {
	t.Helper()
	a, _, dir := built(t)
	m, _ := CurrentManifest(dir)
	fn(dir, m)
	ix, err := Open(a, dir)
	if err == nil {
		ix.Close()
	}
	return err
}
