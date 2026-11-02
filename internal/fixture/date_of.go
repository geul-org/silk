//ff:func feature=fixture type=encoder control=selection
//ff:what dateOf — a random date value: a GREGORIAN-DATE, an EDTF decade or century, or somevalue
package fixture

import (
	"encoding/binary"
	"math/rand"
	"strconv"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/pkg/grammar"
)

func dateOf(r *rand.Rand) []byte {
	switch r.Intn(4) {
	case 0:
		return []byte{0, 1}
	case 1:
		return grammar.RecordOf(core.EncodeBytesValue(core.CodecEDTFEven, []byte("19"+strconv.Itoa(r.Intn(10))+"X")))
	case 2:
		return grammar.RecordOf(core.EncodeBytesValue(core.CodecEDTFEven, []byte("19XX")))
	}
	days := int32(r.Intn(60000) - 40000)
	ws, _ := core.ReadWords(binary.BigEndian.AppendUint32(nil, uint32(days)))
	return grammar.RecordOf(core.Record{Kind: core.KindValue, Codec: core.CodecGregorianDate, Payload: ws})
}
