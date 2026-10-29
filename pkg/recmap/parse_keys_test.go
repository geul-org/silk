//ff:func feature=recmap type=decoder control=iteration dimension=1
//ff:what TestParseKeys — contiguous ascending ranges parse; bad key, padding, start, count, order and coverage fail
package recmap

import (
	"encoding/binary"
	"errors"
	"testing"
)

func TestParseKeys(t *testing.T) {
	enc := func(ks ...KeyRange) []byte { return encodeBody(&Map{Keys: ks}) }
	ok := enc(KeyRange{Key: 1, Start: 0, Count: 2}, KeyRange{Key: 3, Start: 2, Count: 1})
	if ks, err := parseKeys(ok, 3); err != nil || len(ks) != 2 || ks[1].Start != 2 {
		t.Fatal(ks, err)
	}
	pad := enc(KeyRange{Key: 1, Count: 1})
	binary.LittleEndian.PutUint16(pad[4:], 1)
	cases := [][]byte{
		enc(KeyRange{Key: keySpace, Count: 1}),
		pad,
		enc(KeyRange{Key: 1, Start: 1, Count: 1}),
		enc(KeyRange{Key: 1, Count: 0}),
		enc(KeyRange{Key: 1, Count: 9}),
		enc(KeyRange{Key: 3, Count: 1}, KeyRange{Key: 1, Start: 1, Count: 1}),
	}
	for i, c := range cases {
		if _, err := parseKeys(c, 2); !errors.Is(err, ErrFormat) {
			t.Fatal(i, err)
		}
	}
	if _, err := parseKeys(ok, 4); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
