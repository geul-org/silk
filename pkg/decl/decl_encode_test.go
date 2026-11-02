//ff:func feature=decl type=scanner control=sequence
//ff:what TestDeclEncode — JSON keeps every field set
package decl

import (
	"encoding/json"
	"testing"
)

func TestDeclEncode(t *testing.T) {
	var d Decl
	b := Decl{Name: "n", Kind: KindText, Tags: []uint16{1}, Keys: []string{"k"}, Widths: []int{8}}.Encode()
	if err := json.Unmarshal(b, &d); err != nil || d.Name != "n" || d.Tags[0] != 1 || d.Widths[0] != 8 {
		t.Fatal(string(b), err)
	}
}
