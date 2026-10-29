//ff:func feature=seg type=reader control=sequence
//ff:what TestFileEach — every row once with its partition key; an fn error, a group read error and a row read error stop
package seg

import (
	"errors"
	"io"
	"testing"
)

func TestFileEach(t *testing.T) {
	f := sampleSeg(t, true)
	keys := ""
	if err := f.Each(func(p, r []byte) error { keys += string(p); return nil }); err != nil || keys != "aaaaaaccccbbdd" {
		t.Fatal(err, keys)
	}
	if err := f.Each(func(p, r []byte) error { return io.EOF }); !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
	if err := failing(f, 0).Each(func(p, r []byte) error { return nil }); err == nil {
		t.Fatal("groups")
	}
	if err := failing(f, 1).Each(func(p, r []byte) error { return nil }); err == nil {
		t.Fatal("rows")
	}
}
