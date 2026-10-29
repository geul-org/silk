//ff:func feature=seg type=rule control=sequence
//ff:what TestFileVerify — written segments verify; a read error, a changed body byte, rows out of order and a key in both places fail
package seg

import (
	"bytes"
	"testing"
)

func TestFileVerify(t *testing.T) {
	f := sampleSeg(t, true)
	if err := f.Verify(); err != nil {
		t.Fatal(err)
	}
	if err := failing(f, 0).Verify(); err == nil {
		t.Fatal("read error")
	}
	raw := make([]byte, f.Info.Bytes)
	f.r.ReadAt(raw, 0)
	flip := append([]byte{}, raw...)
	flip[0] ^= 1
	g := *f
	g.r = bytes.NewReader(flip)
	if err := g.Verify(); err == nil {
		t.Fatal("hash")
	}
	if err := resealed(f, swapRows(raw, 0, 10, 10)).Verify(); err == nil {
		t.Fatal("partition rows out of order")
	}
	if err := resealed(f, swapRows(raw, 50, 62, 12)).Verify(); err == nil {
		t.Fatal("other rows out of order")
	}
	dup := append([]byte{}, raw...)
	copy(dup[50:52], "aa")
	copy(dup[62:64], "aa")
	if err := resealed(f, dup).Verify(); err == nil {
		t.Fatal("key in the table and in the other partition")
	}
}
