//ff:func feature=seg type=decoder control=sequence
//ff:what TestSeqOf — the document and offset at the end of a row
package seg

import "testing"

func TestSeqOf(t *testing.T) {
	if d, o := SeqOf(row("ab", 7, 300)); d != 7 || o != 300 {
		t.Fatal(d, o)
	}
}
