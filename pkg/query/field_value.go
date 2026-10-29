//ff:func feature=query type=decoder control=sequence
//ff:what fieldValue — the bytes at a field path of decoded fields: the field record, or the first record of its payload (a language record); false when absent
package query

import "github.com/geul-org/gogeul/pkg/core"

// fieldValue reads path with gogeul core.
func fieldValue(fs [][]core.Word, path []int) ([]byte, bool) {
	f := fs[path[0]]
	if len(path) == 1 {
		return core.PutWords(f), true
	}
	r, _, err := core.DecodeRecord(f)
	if err != nil || r.Kind != core.KindValue || r.Codec&^1 != core.CodecUTF8LanguageTextEven {
		return nil, false
	}
	_, n, err := core.DecodeRecord(r.Payload)
	if err != nil {
		return nil, false
	}
	return core.PutWords(r.Payload[:n]), true
}
