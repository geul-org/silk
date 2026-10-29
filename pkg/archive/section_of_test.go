//ff:func feature=archive type=decoder control=sequence
//ff:what TestSectionOf — finds a section past others; truncated tables, out-of-bounds and missing sections fail
package archive

import (
	"encoding/binary"
	"errors"
	"testing"
)

func TestSectionOf(t *testing.T) {
	b := make([]byte, manifestHead+2*sectionSize)
	entry := func(i int, kind uint16, off, n uint64) {
		e := b[manifestHead+i*sectionSize:]
		binary.LittleEndian.PutUint16(e, kind)
		binary.LittleEndian.PutUint64(e[8:], off)
		binary.LittleEndian.PutUint64(e[16:], n)
	}
	entry(0, 5, 0, 1)
	entry(1, secObjects, 2, 3)
	if sec, err := sectionOf(b, 2, secObjects); err != nil || len(sec) != 3 {
		t.Fatal(sec, err)
	}
	if _, err := sectionOf(b, 3, secObjects); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	if _, err := sectionOf(b, 1, secObjects); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	entry(1, secObjects, 2, uint64(len(b)))
	if _, err := sectionOf(b, 2, secObjects); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
}
