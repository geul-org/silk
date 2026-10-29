//ff:func feature=seg type=reader control=sequence
//ff:what TestEachRow — each row of a group with the key; fn and read errors stop
package seg

import (
	"errors"
	"io"
	"testing"
)

func TestEachRow(t *testing.T) {
	f := sampleSeg(t, true)
	gs, _ := f.Groups()
	n := 0
	if err := eachRow(gs[0], func(p, r []byte) error { n++; return nil }); err != nil || n != 3 {
		t.Fatal(err, n)
	}
	if err := eachRow(gs[0], func(p, r []byte) error { return io.EOF }); !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
	if err := eachRow(Group{f: failing(f, 0), rg: gs[0].rg}, func(p, r []byte) error { return nil }); err == nil {
		t.Fatal("read error")
	}
}
