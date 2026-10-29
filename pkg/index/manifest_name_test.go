//ff:func feature=index type=helper control=sequence
//ff:what TestManifestName — manifest/<16 hex digits>.json
package index

import "testing"

func TestManifestName(t *testing.T) {
	if n := manifestName(26); n != "manifest/000000000000001a.json" {
		t.Fatal(n)
	}
}
