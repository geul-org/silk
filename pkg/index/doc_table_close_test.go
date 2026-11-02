//ff:func feature=index type=scanner control=sequence
//ff:what TestDocTableClose — a table unmaps its files
package index

import "testing"

func TestDocTableClose(t *testing.T) {
	_, dir, _ := built(t, 1)
	m := current(t, dir)
	tb, err := openDocs(dir, &m)
	if err != nil || len(tb.files) != 1 {
		t.Fatal(err)
	}
	tb.Close()
	(&DocTable{}).Close()
}
