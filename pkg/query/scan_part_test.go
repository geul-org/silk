//ff:func feature=query type=scanner control=sequence
//ff:what TestScanPart — hits land in their documents with Edge Ids; off[] read, scan, document lookup and exact check failures surface
package query

import (
	"errors"
	"io"
	"testing"

	"github.com/park-jun-woo/silk/pkg/scan"
)

func TestScanPart(t *testing.T) {
	ix := fixtureIndex(t)
	var st Stats
	out := map[docKey][]uint64{}
	phys := mustTriple(t, "P106=Q169470")
	if err := scanPart(ix.Parts[0], phys, scan.Options{}, &st, out); err != nil || len(out) != 2 || st.Hits != 2 {
		t.Fatal(out, err)
	}
	rec := tripleBytes(4, 106, 7, 169470)
	if err := scanPart(craftPart(t, rec, true, nil), phys, scan.Options{}, &st, out); err != nil || len(out[docKey{"0x0001", 9, 1}]) != 1 {
		t.Fatal(out, err)
	}
	tail := func(b []byte) io.ReaderAt { return tailFail{b: b, from: int64(len(b)) - 4} }
	if err := scanPart(craftPart(t, rec, true, tail), phys, scan.Options{}, &st, out); err == nil {
		t.Fatal("off[] read must fail")
	}
	if err := scanPart(craftPart(t, rec[:5], true, nil), phys, scan.Options{}, &st, out); !errors.Is(err, scan.ErrRecord) {
		t.Fatal(err)
	}
	if err := scanPart(craftPart(t, rec, false, nil), phys, scan.Options{}, &st, out); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
	long := append([]byte{rec[0] | 0x07, rec[1] &^ 0x80, 0, byte(len(rec)/2 + 1)}, rec[2:]...)
	if err := scanPart(craftPart(t, long, true, nil), phys, scan.Options{}, &st, out); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
}
