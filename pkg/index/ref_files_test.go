//ff:func feature=index type=helper control=sequence
//ff:what TestRefFiles — file names in order
package index

import (
	"slices"
	"testing"
)

func TestRefFiles(t *testing.T) {
	if f := refFiles([]SegRef{{File: "b"}, {File: "a"}}); !slices.Equal(f, []string{"b", "a"}) {
		t.Fatal(f)
	}
}
