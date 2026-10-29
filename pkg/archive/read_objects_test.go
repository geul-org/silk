//ff:func feature=archive type=decoder control=sequence
//ff:what TestReadObjects — entries decode kind, size and hash hex; truncated and wrong-length bodies fail
package archive

import (
	"encoding/binary"
	"errors"
	"strings"
	"testing"
)

func TestReadObjects(t *testing.T) {
	sec := make([]byte, 4+objectSize)
	binary.LittleEndian.PutUint32(sec, 1)
	sec[4] = kindData
	binary.LittleEndian.PutUint64(sec[12:], 77)
	sec[20] = 0xab
	objs, err := readObjects(sec)
	if err != nil || len(objs) != 1 || objs[0].Kind != kindData || objs[0].Size != 77 || !strings.HasPrefix(objs[0].Hash, "ab") {
		t.Fatal(objs, err)
	}
	if _, err := readObjects(sec[:3]); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
	if _, err := readObjects(sec[:10]); !errors.Is(err, ErrManifest) {
		t.Fatal(err)
	}
}
