//ff:func feature=recmap type=rule control=sequence
//ff:what TestCheckDocs — ascending documents inside the pack pass; huge packs, overlaps and documents past the end fail
package recmap

import (
	"errors"
	"math"
	"testing"

	"github.com/park-jun-woo/silk/pkg/archive"
)

func TestCheckDocs(t *testing.T) {
	ok := []archive.Doc{{Off: 64, Len: 10}, {Off: 74, Len: 6}}
	if err := checkDocs(ok, 80); err != nil {
		t.Fatal(err)
	}
	if err := checkDocs(nil, math.MaxUint32+1); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := checkDocs([]archive.Doc{{Off: 64, Len: 10}, {Off: 70, Len: 2}}, 80); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := checkDocs([]archive.Doc{{Off: 75, Len: 10}}, 80); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := checkDocs([]archive.Doc{{Off: 0, Len: 90}}, 80); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
