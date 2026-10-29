//ff:func feature=seg type=rule control=sequence
//ff:what TestFileCheckOther — other keys absent from the table pass; a key also in the table or a read error fails
package seg

import "testing"

func TestFileCheckOther(t *testing.T) {
	f := sampleSeg(t, true)
	if err := f.checkOther(); err != nil {
		t.Fatal(err)
	}
	raw := make([]byte, f.Info.Bytes)
	f.r.ReadAt(raw, 0)
	copy(raw[50:52], "aa")
	if err := resealed(f, raw).checkOther(); err == nil {
		t.Fatal("key in the table and in the other partition")
	}
	if err := failing(f, 0).checkOther(); err == nil {
		t.Fatal("read error")
	}
}
