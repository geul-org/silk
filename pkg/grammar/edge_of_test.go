//ff:func feature=grammar type=helper control=sequence
//ff:what edgeOf — an Edge record from field bytes
package grammar

import (
	"bytes"

	"github.com/geul-org/gogeul/pkg/core"
)

func edgeOf(codec uint8, fs ...[]byte) []byte {
	ws, _ := core.ReadWords(bytes.Join(fs, nil))
	return RecordOf(core.Record{Kind: core.KindEdge, Codec: codec, Payload: ws})
}
