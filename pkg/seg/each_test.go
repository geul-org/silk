//ff:func feature=seg type=scanner control=sequence
//ff:what TestEach — rows come out merged; Live drops, a callback error stops
package seg

import (
	"encoding/binary"
	"errors"
	"slices"
	"testing"
)

func TestEach(t *testing.T) {
	a, b := writeSeg(t, "a"), writeSeg(t, "b")
	var docs []uint32
	err := Each([]*File{a, b}, Hooks{Live: func(d uint32) bool { return d != 3 }}, func(part, r []byte) error {
		docs = append(docs, binary.BigEndian.Uint32(r[2:]))
		return nil
	})
	if err != nil || !slices.Equal(docs, []uint32{1, 1, 2, 2}) {
		t.Fatal(docs, err)
	}
	boom := errors.New("stop")
	if err := Each([]*File{a}, Hooks{}, func(_, _ []byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
	if err := visit(&cursor{f: a, reg: a.All()}, Hooks{}, func(_, _ []byte) error { return nil }); err != nil {
		t.Fatal(err)
	}
}
