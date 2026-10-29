//ff:func feature=seg type=builder control=sequence
//ff:what TestBucketAdd — rows gathered under their partition key, counted
package seg

import "testing"

func TestBucketAdd(t *testing.T) {
	b := NewBucket(Info{}, 1)
	b.Add([]byte("aa"), row("k1", 1, 0))
	b.Add([]byte("aa"), row("k2", 1, 0))
	if b.Rows != 2 || len(b.parts["aa"]) != 20 {
		t.Fatal(b.Rows)
	}
}
