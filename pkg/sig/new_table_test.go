//ff:func feature=sig type=builder control=sequence
//ff:what TestNewTable — an empty table as wide as the schema
package sig

import "testing"

func TestNewTable(t *testing.T) {
	if tb := NewTable(sampleSchema()); tb.Words != 2 || tb.Len() != 0 {
		t.Fatal(tb)
	}
}
