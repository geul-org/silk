//ff:func feature=index type=helper control=sequence
//ff:what TestFill — n copies of a value
package index

import (
	"slices"
	"testing"
)

func TestFill(t *testing.T) {
	if !slices.Equal(fill(3, true), []bool{true, true, true}) || len(fill(0, true)) != 0 {
		t.Fatal("fill")
	}
}
