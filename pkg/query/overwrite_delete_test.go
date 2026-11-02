//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what TestOverwriteDelete — a rewritten document's old rows and a deleted document's rows never answer, by index or by scan
package query

import (
	"context"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func TestOverwriteDelete(t *testing.T) {
	ctx := context.Background()
	s := fixture.Tok("Q7")
	p := fixture.Tok("P106")
	v1 := fixture.Doc{ID: 0x5100000000000007, Version: 1, Body: []core.Record{fixture.Triple(1, p, s, fixture.Tok("Q82955"), 0), fixture.Label(2, s, "ko", "옛 이름")}}
	v2 := fixture.Doc{ID: v1.ID, Version: 1, Body: []core.Record{fixture.Triple(1, p, s, fixture.Tok("Q36180"), 0), fixture.Label(2, s, "ko", "새 이름")}}
	gone := fixture.Doc{ID: 0x5100000000000008, Version: 1, Body: []core.Record{fixture.Triple(1, p, fixture.Tok("Q8"), fixture.Tok("Q82955"), 0)}}
	data := fixture.Pack(t, []fixture.Doc{v1, gone}, nil)
	a, dir, ix := openIndex(t, data)
	ix.Close()
	fixture.Repack(t, data, []fixture.Doc{v2}, nil)
	a, _ = archive.Open(ctx, data)
	if _, err := index.Update(ctx, a, dir, testOpts); err != nil {
		t.Fatal(err)
	}
	ix2, err := index.Open(ctx, a, dir, 0)
	if err != nil {
		t.Fatal(err)
	}
	defer ix2.Close()
	for _, scan := range []bool{false, true} {
		if n := answer(t, ix2, scan, "TRIPLE", "predicate=P106", "object=Q82955"); len(n) != 0 {
			t.Fatal(scan, n)
		}
		if n := answer(t, ix2, scan, "LABEL", `object.bytes="옛 이름"`); len(n) != 0 {
			t.Fatal(scan, n)
		}
		if n := answer(t, ix2, scan, "TRIPLE", "predicate=P106", "object=Q36180"); len(n) != 1 {
			t.Fatal(scan, n)
		}
	}
	if ix2.Live.Get(0) || ix2.Live.Get(1) {
		t.Fatal("old numbers live")
	}
}
