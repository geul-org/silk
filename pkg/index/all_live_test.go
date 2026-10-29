//ff:func feature=index type=helper control=sequence
//ff:what TestAllLive — n live sequence numbers
package index

import "testing"

func TestAllLive(t *testing.T) {
	b := allLive(10)
	if b.N != 10 || b.Count() != 10 || b.Get(10) {
		t.Fatal(b)
	}
}
