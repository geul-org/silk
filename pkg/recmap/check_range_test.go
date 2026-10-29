//ff:func feature=recmap type=helper control=iteration dimension=1
//ff:what checkRange — every offset of a key range starts a record whose TAG WORD has that Kind·Codec
package recmap

import (
	"encoding/binary"
	"io"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkRange reads each record's TAG WORD from data.
func checkRange(t *testing.T, data io.ReaderAt, k KeyRange, offs []uint32) {
	t.Helper()
	for _, off := range offs {
		var tag [2]byte
		data.ReadAt(tag[:], int64(off))
		kind, codec := core.SplitTarget(core.Word(binary.BigEndian.Uint16(tag[:])))
		if uint16(core.TargetWord(kind, codec)) != k.Key {
			t.Fatal(k, off)
		}
	}
}
