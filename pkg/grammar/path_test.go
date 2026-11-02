//ff:func feature=grammar type=scanner control=sequence
//ff:what TestPath — numbers join with dots
package grammar

import "testing"

func TestPath(t *testing.T) {
	if s := (Path{2, 1}).String(); s != "2.1" {
		t.Fatal(s)
	}
}
