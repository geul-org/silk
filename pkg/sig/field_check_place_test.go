//ff:func feature=sig type=scanner control=sequence
//ff:what TestFieldCheckPlace — bits marked, shared ones refused, a wide flag refused
package sig

import "testing"

func TestFieldCheckPlace(t *testing.T) {
	used := make([]bool, 4)
	f := &Field{Off: 1, Width: 2}
	if err := f.checkPlace(used); err != nil || !used[2] {
		t.Fatal(err)
	}
	if err := f.checkPlace(used); err == nil {
		t.Fatal("shared")
	}
	if err := (&Field{Width: 9}).checkPlace(used); err == nil {
		t.Fatal("past the end")
	}
	if err := (&Field{Width: 2, Kind: KindFlag}).checkPlace(used); err == nil {
		t.Fatal("wide flag")
	}
}
