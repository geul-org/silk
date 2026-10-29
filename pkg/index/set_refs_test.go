//ff:func feature=index type=helper control=sequence
//ff:what TestSetRefs — the archive's sets with names, generations and manifests
package index

import "testing"

func TestSetRefs(t *testing.T) {
	a, _, _ := built(t)
	s := setRefs(a)
	if len(s) != len(a.Sets) || s[0].Manifest != a.Sets[0].Manifest || s[0].Generation != 1 {
		t.Fatal(s)
	}
}
