//ff:func feature=query type=helper control=sequence
//ff:what tripleBytes — a TRIPLE record (Edge Id #id, P pred, Q subj, Q obj, rank normal) encoded with gogeul core
package query

import "github.com/geul-org/gogeul/pkg/core"

// tripleBytes builds the record bytes.
func tripleBytes(id, pred, subj, obj uint64) []byte {
	ref := func(c uint8, n uint64) []core.Word {
		ws, _ := core.EncodeRecord(core.Record{Kind: core.KindReference, Codec: c << 3, Payload: core.EncodeUintValue(n).Payload})
		return ws
	}
	pay := append(append(append(append(ref(core.RefLocal, id), ref(core.RefWikidataPID, pred)...), ref(core.RefWikidataQID, subj)...), ref(core.RefWikidataQID, obj)...), 0)
	ws, _ := core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: core.EdgeTriple, Payload: pay})
	return core.PutWords(ws)
}
