//ff:func feature=scan type=matcher control=sequence
//ff:what TestTripleMatch — Predicate/Object equality, wildcards, long (Length extension) records and truncated fields
package scan

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestTripleMatch(t *testing.T) {
	rec, p, o := tripleRec(3, 106, 42, 169470)
	_, p2, o2 := tripleRec(3, 31, 42, 5)
	if !(Triple{Pred: p, Obj: o}).Match(rec) || !(Triple{Pred: p}).Match(rec) || !(Triple{Obj: o}).Match(rec) {
		t.Fatal("must match")
	}
	if (Triple{Pred: p2, Obj: o}).Match(rec) || (Triple{Pred: p, Obj: o2}).Match(rec) {
		t.Fatal("must not match")
	}
	if (Triple{}).Match(rec[:4]) || (Triple{}).Match(rec[:len(rec)-6]) {
		t.Fatal("truncated must not match")
	}
	text := core.EncodeBytesValue(core.CodecUTF8TextEven, make([]byte, 40))
	tw, _ := core.EncodeRecord(text)
	pay := append(append(append(append([]core.Word{0x0088, 0x0001}, 0x00a8, 0x006a), 0x00a0, 0x002a), tw...), 0)
	ws, _ := core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: core.EdgeTriple, Payload: pay})
	if !(Triple{Pred: p}).Match(core.PutWords(ws)) {
		t.Fatal("long TRIPLE")
	}
}
