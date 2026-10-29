//ff:func feature=spec type=encoder control=sequence
//ff:what TestTagExtract — TRIPLE keys (Predicate slot, Object slot), a tag without partition, and zero slots for a record missing the field
package spec

import (
	"bytes"
	"testing"
)

func TestTagExtract(t *testing.T) {
	tr := hexBytes("148a 0088 0005 00a8 001f 00a0 002a 00a0 0005 0000")
	g := Default().Tags[0]
	if k := g.Extract(nil, tr); !bytes.Equal(k, hexBytes("00a8 001f 0000 00a0 0005 0000")) {
		t.Fatalf("%x", k)
	}
	src := Tag{Sort: Key{Path: []int{2}, Width: 4}}
	if k := src.Extract(nil, hexBytes("1320 0088 00bb 0088 0000 0088 00b8")); !bytes.Equal(k, hexBytes("0088 00b8")) {
		t.Fatalf("%x", k)
	}
	if k := g.Extract(nil, hexBytes("108a 0088 0005")); !bytes.Equal(k, make([]byte, 12)) {
		t.Fatalf("%x", k)
	}
}
