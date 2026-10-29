//ff:func feature=index type=helper control=sequence
//ff:what TestSetOf — a set is found by manifest; an unknown manifest is not
package index

import (
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestSetOf(t *testing.T) {
	a := &archive.Archive{Sets: []archive.Set{{Name: "a", Manifest: "m1"}, {Name: "b", Manifest: "m2"}}}
	if s, ok := setOf(a, "m2"); !ok || s.Name != "b" {
		t.Fatal(s, ok)
	}
	if _, ok := setOf(a, "m3"); ok {
		t.Fatal("m3")
	}
}
