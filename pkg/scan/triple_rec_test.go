//ff:func feature=scan type=helper control=sequence
//ff:what tripleRec — the bytes of a TRIPLE record (Edge Id #id, P pred, Q subj, Q obj, rank normal) built with gogeul core
package scan

import "github.com/geul-org/gogeul/pkg/core"

// tripleRec encodes one TRIPLE and returns it with its Predicate and Object record bytes.
func tripleRec(id, pred, subj, obj uint64) (rec, p, o []byte) {
	ref := func(c uint8, n uint64) []core.Word {
		ws, _ := core.EncodeRecord(core.Record{Kind: core.KindReference, Codec: c << 3, Payload: core.EncodeUintValue(n).Payload})
		return ws
	}
	pw, ow := ref(core.RefWikidataPID, pred), ref(core.RefWikidataQID, obj)
	pay := append(append(append(append(ref(core.RefLocal, id), pw...), ref(core.RefWikidataQID, subj)...), ow...), 0)
	ws, _ := core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: core.EdgeTriple, Payload: pay})
	return core.PutWords(ws), core.PutWords(pw), core.PutWords(ow)
}
