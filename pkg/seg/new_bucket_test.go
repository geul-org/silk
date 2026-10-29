//ff:func feature=seg type=builder control=sequence
//ff:what TestNewBucket — an empty bucket keeping its Info and minimum
package seg

import "testing"

func TestNewBucket(t *testing.T) {
	b := NewBucket(Info{Tag: 7}, 3)
	if b.Info.Tag != 7 || b.Min != 3 || b.Rows != 0 || len(b.parts) != 0 {
		t.Fatal(b)
	}
}
