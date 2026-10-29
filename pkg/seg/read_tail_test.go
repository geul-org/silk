//ff:func feature=seg type=reader control=sequence
//ff:what TestReadTail — a good tail parses; a short file, failing reads, a bad footer, a size mismatch or a bad table fail
package seg

import (
	"bytes"
	"testing"
)

func TestReadTail(t *testing.T) {
	f := sampleSeg(t, true)
	size := f.Info.Bytes
	if g, err := readTail(f.r, size); err != nil || g.Info.Parts != 2 {
		t.Fatal(err)
	}
	if _, err := readTail(f.r, 10); err == nil {
		t.Fatal("short")
	}
	if _, err := readTail(&failReader{r: f.r}, size); err == nil {
		t.Fatal("footer read")
	}
	if _, err := readTail(&failReader{r: f.r, n: 1}, size); err == nil {
		t.Fatal("table read")
	}
	raw := make([]byte, size)
	f.r.ReadAt(raw, 0)
	nomagic := append([]byte{}, raw...)
	nomagic[size-1] = 'X'
	if _, err := readTail(bytes.NewReader(nomagic), size); err == nil {
		t.Fatal("bad footer")
	}
	big := append(append([]byte{}, raw[:size-footerSize]...), append(make([]byte, 18), raw[size-footerSize:]...)...)
	if _, err := readTail(bytes.NewReader(big), int64(len(big))); err == nil {
		t.Fatal("size mismatch")
	}
	tbl := size - footerSize - 2*18
	bad := append([]byte{}, raw...)
	bad[tbl+2+7] = 5
	if _, err := readTail(bytes.NewReader(bad), size); err == nil {
		t.Fatal("bad table")
	}
}
