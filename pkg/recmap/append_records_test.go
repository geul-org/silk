//ff:func feature=recmap type=builder control=sequence
//ff:what TestAppendRecords — a DOC/BODY/value/END document gives four keys at word offsets from base; odd bytes and non-documents fail
package recmap

import (
	"errors"
	"slices"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestAppendRecords(t *testing.T) {
	doc, err := core.EncodeDocument(nil, []core.Record{core.EncodeUintValue(5)})
	if err != nil {
		t.Fatal(err)
	}
	keys, offs, err := appendRecords([]uint16{9}, []uint32{1}, doc, 100)
	meta := func(c uint8) uint16 { return uint16(core.TargetWord(core.KindMeta, c)) }
	wantKeys := []uint16{9, meta(core.MetaDoc1Word), meta(core.MetaBody), uint16(core.TargetWord(core.KindValue, core.CodecUnsignedInteger)), meta(core.MetaEnd)}
	if err != nil || !slices.Equal(keys, wantKeys) || !slices.Equal(offs, []uint32{1, 100, 104, 106, 110}) {
		t.Fatal(keys, offs, err)
	}
	if _, _, err := appendRecords(nil, nil, doc[:3], 0); !errors.Is(err, ErrDocument) {
		t.Fatal(err)
	}
	if _, _, err := appendRecords(nil, nil, doc[:4], 0); !errors.Is(err, ErrDocument) {
		t.Fatal(err)
	}
}
