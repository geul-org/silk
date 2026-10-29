//ff:func feature=index type=helper control=sequence
//ff:what TestIndexSection — the section and tag of TRIPLE; false for a kind not indexed
package index

import "testing"

func TestIndexSection(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	if i, tg, ok := ix.Section(0x100A); !ok || i != 0 || tg.Name != "TRIPLE" {
		t.Fatal(i, tg, ok)
	}
	if _, _, ok := ix.Section(0x0801); ok {
		t.Fatal("Value kind")
	}
}
