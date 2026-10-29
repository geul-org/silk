//ff:func feature=seg type=reader control=sequence
//ff:what TestFileEmitRange — every row of a range in chunks; a read error or an emit error stops it
package seg

import (
	"errors"
	"io"
	"slices"
	"testing"
)

func TestFileEmitRange(t *testing.T) {
	f := sampleSeg(t, true)
	emit, got := collect()
	if err := f.emitRange(f.partRegion(0), 0, 3, nil, 0, false, emit); err != nil || !slices.Equal(*got, []string{"1:4", "2:4", "3:9"}) {
		t.Fatal(err, *got)
	}
	if err := failing(f, 0).emitRange(f.partRegion(0), 0, 3, nil, 0, false, emit); err == nil {
		t.Fatal("read error")
	}
	stop := func(uint32, uint32) error { return io.EOF }
	if err := f.emitRange(f.partRegion(0), 0, 3, nil, 0, false, stop); !errors.Is(err, io.EOF) {
		t.Fatal(err)
	}
}
