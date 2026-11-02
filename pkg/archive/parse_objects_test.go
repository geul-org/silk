//ff:func feature=archive type=scanner control=sequence
//ff:what TestParseObjects — headers, section tables and OBJECTS bodies are bounds checked
package archive

import (
	"encoding/binary"
	"errors"
	"slices"
	"testing"
)

func TestParseObjects(t *testing.T) {
	if _, err := parseObjects([]byte("x")); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	head := append([]byte(manifestMagic), make([]byte, manifestHead-8)...)
	binary.LittleEndian.PutUint32(head[12:], 1)
	if _, err := parseObjects(head); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	e := make([]byte, sectionSize)
	binary.LittleEndian.PutUint16(e, secObjects)
	binary.LittleEndian.PutUint64(e[8:], manifestHead+sectionSize)
	binary.LittleEndian.PutUint64(e[16:], 52)
	body := make([]byte, 52)
	binary.LittleEndian.PutUint32(body, 1)
	body[4] = 1
	objs, err := parseObjects(append(append(slices.Clone(head), e...), body...))
	if err != nil || len(objs) != 1 || objs[0].Kind != 1 {
		t.Fatal(objs, err)
	}
	if _, err := sectionOf(append(slices.Clone(head), e...), 1, secObjects); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	if _, err := sectionOf(append(slices.Clone(head), e...), 1, 0x9); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	if _, err := sectionOf(head, -1, 1); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	if _, err := readObjects([]byte{1}); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	if _, err := readObjects([]byte{1, 0, 0, 0}); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
}
