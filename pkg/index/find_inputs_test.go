//ff:func feature=index type=helper control=iteration dimension=1
//ff:what TestFindInputs — the section and references when every file is still there; a missing file, no files or an unknown section are false
package index

import "testing"

func TestFindInputs(t *testing.T) {
	secs := []Section{{Name: "A"}, {Name: "B", Runs: []SegRef{{File: "r"}}, Appends: []SegRef{{File: "p"}}}}
	if i, rs, ok := findInputs(secs, "B", []string{"p", "r"}); !ok || i != 1 || rs[0].File != "p" {
		t.Fatal(i, rs, ok)
	}
	for _, c := range [][]string{{"r", "x"}, {}} {
		if _, _, ok := findInputs(secs, "B", c); ok {
			t.Fatal(c)
		}
	}
	if _, _, ok := findInputs(secs, "C", []string{"r"}); ok {
		t.Fatal("unknown section")
	}
}
