//ff:func feature=index type=helper control=sequence
//ff:what TestFileSet — every reference's file
package index

import "testing"

func TestFileSet(t *testing.T) {
	s := fileSet([]SegRef{{File: "a"}, {File: "b"}})
	if len(s) != 2 || !s["a"] || !s["b"] {
		t.Fatal(s)
	}
}
