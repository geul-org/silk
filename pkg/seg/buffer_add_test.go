//ff:func feature=seg type=scanner control=sequence
//ff:what TestBufferAdd — partition and row sit side by side
package seg

import "testing"

func TestBufferAdd(t *testing.T) {
	b := NewBuffer(testInfo())
	b.Add([]byte("pa"), row(0, 1, 1))
	if b.Rows() != 1 || string(b.at(0)[:2]) != "pa" {
		t.Fatal(b.Rows())
	}
}
