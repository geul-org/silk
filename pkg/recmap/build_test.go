//ff:func feature=recmap type=builder control=iteration dimension=1
//ff:what TestBuild — every offset of the fixture map frames a record of its range's key; bad order, unreadable and non-GEUL documents fail
package recmap

import (
	"errors"
	"math"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestBuild(t *testing.T) {
	p, f := fixturePack(t)
	m, err := Build(f, p.Size, p.Docs)
	if err != nil || len(m.Docs) != 4 || m.Header.Records != uint64(len(m.Offs)) || m.Header.Keys != uint32(len(m.Keys)) {
		t.Fatal(m, err)
	}
	for _, k := range m.Keys {
		checkRange(t, f, k, m.Offs[k.Start:k.Start+k.Count])
	}
	if _, err := Build(f, math.MaxUint32+1, nil); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if _, err := Build(failReader{from: 0}, p.Size, p.Docs); err == nil {
		t.Fatal("read error must fail")
	}
	if _, err := Build(f, p.Size, []archive.Doc{{Off: 0, Len: 64}}); !errors.Is(err, ErrDocument) {
		t.Fatal(err)
	}
}
