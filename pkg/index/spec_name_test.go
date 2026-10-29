//ff:func feature=index type=helper control=sequence
//ff:what TestSpecName — spec/<hash>.json
package index

import "testing"

func TestSpecName(t *testing.T) {
	if specName("ab") != "spec/ab.json" {
		t.Fatal(specName("ab"))
	}
}
