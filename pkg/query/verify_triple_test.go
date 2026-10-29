//ff:func feature=query type=matcher control=iteration dimension=1
//ff:what TestVerifyTriple — the Edge Id of a matching TRIPLE; odd bytes, other records, short payloads, other values and non-LOCAL Edge Ids fail
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestVerifyTriple(t *testing.T) {
	phys := mustTriple(t, "P106=Q169470")
	rec := tripleBytes(12, 106, 7, 169470)
	if id, err := verifyTriple(rec, phys); err != nil || id != 12 {
		t.Fatal(id, err)
	}
	if id, err := verifyTriple(rec, Triple{}); err != nil || id != 12 {
		t.Fatal(id, err)
	}
	other := tripleBytes(12, 31, 7, 5)
	label := core.PutWords(must(core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: 0x0C, Payload: []core.Word{0x0088, 1}})))
	short := core.PutWords(must(core.EncodeRecord(core.Record{Kind: core.KindEdge, Codec: core.EdgeTriple, Payload: []core.Word{0x0088, 1}})))
	notLocal := append([]byte(nil), rec...)
	notLocal[3] = 0xa0
	for i, b := range [][]byte{rec[:3], label, short, other, notLocal} {
		if _, err := verifyTriple(b, phys); !errors.Is(err, ErrVerify) {
			t.Fatal(i, err)
		}
	}
}
