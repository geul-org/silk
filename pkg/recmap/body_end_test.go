//ff:func feature=recmap type=helper control=sequence
//ff:what TestBodyEnd — header + 24K + 24D + 4N; overflowing document or record counts saturate
package recmap

import (
	"math"
	"testing"
)

func TestBodyEnd(t *testing.T) {
	if got := bodyEnd(Header{Keys: 1, Docs: 2, Records: 3}); got != headerSize+24+48+12 {
		t.Fatal(got)
	}
	if got := bodyEnd(Header{Docs: math.MaxUint64 / 2}); got != math.MaxUint64 {
		t.Fatal(got)
	}
	if got := bodyEnd(Header{Records: math.MaxUint64/4 - 10, Docs: 1}); got != math.MaxUint64 {
		t.Fatal(got)
	}
}
