//ff:func feature=index type=helper control=sequence
//ff:what TestSectionNamed — the section of a name, a zero Section otherwise
package index

import "testing"

func TestSectionNamed(t *testing.T) {
	s := []Section{{Name: "A", Dead: 1}}
	if sectionNamed(s, "A").Dead != 1 || sectionNamed(s, "B").Name != "" {
		t.Fatal("sectionNamed")
	}
}
