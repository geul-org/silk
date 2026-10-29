//ff:func feature=seg type=builder control=sequence
//ff:what TestBucketReset — rows dropped, count zero
package seg

import "testing"

func TestBucketReset(t *testing.T) {
	b := NewBucket(Info{}, 1)
	b.Add([]byte("aa"), row("k1", 1, 0))
	b.Reset()
	if b.Rows != 0 || len(b.parts) != 0 {
		t.Fatal(b.Rows)
	}
}
