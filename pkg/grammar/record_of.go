//ff:func feature=grammar type=encoder control=sequence
//ff:what RecordOf — a Reference, Value or other record as bytes
package grammar

import "github.com/geul-org/gogeul/pkg/core"

func RecordOf(r core.Record) []byte {
	ws, _ := core.EncodeRecord(r)
	return core.PutWords(ws)
}
