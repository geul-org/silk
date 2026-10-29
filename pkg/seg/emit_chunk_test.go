//ff:func feature=seg type=matcher control=sequence
//ff:what TestEmitChunk — rows whose bytes at an offset equal the prefix (all without comparison); an emit error stops
package seg

import (
	"errors"
	"io"
	"slices"
	"testing"
)

func TestEmitChunk(t *testing.T) {
	b := append(append([]byte("aax1"), row("", 1, 2)...), append([]byte("bbx2"), row("", 3, 4)...)...)
	emit, got := collect()
	if err := emitChunk(b, 12, []byte("x2"), 2, true, emit); err != nil || !slices.Equal(*got, []string{"3:4"}) {
		t.Fatal(err, *got)
	}
	if err := emitChunk(b, 12, nil, 0, false, func(uint32, uint32) error { return io.EOF }); !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
}
