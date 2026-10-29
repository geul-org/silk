//ff:func feature=sig type=helper control=sequence
//ff:what TestSchemaWords — 16 and 64 bits take one word, 128 two
package sig

import "testing"

func TestSchemaWords(t *testing.T) {
	if (Schema{Width: 16}).Words() != 1 || (Schema{Width: 64}).Words() != 1 || (Schema{Width: 128}).Words() != 2 {
		t.Fatal("words")
	}
}
