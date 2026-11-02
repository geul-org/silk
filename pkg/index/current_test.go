//ff:func feature=index type=helper control=sequence
//ff:what current — the manifest of an index directory
package index

import "testing"

func current(t *testing.T, dir string) Manifest {
	t.Helper()
	m, err := currentManifest(dir)
	if err != nil {
		t.Fatal(err)
	}
	return m
}
