//ff:func feature=query type=matcher control=sequence
//ff:what TestCheckRecord — a matching TRIPLE gives its Edge Id, a differing field false; odd bytes, another kind, a short payload or a bad Edge Id fail
package query

import (
	"errors"
	"testing"
)

func TestCheckRecord(t *testing.T) {
	rec := []byte{0x14, 0x8a, 0, 0x88, 0, 5, 0, 0xa8, 0, 0x1f, 0, 0xa0, 0, 0x2a, 0, 0xa0, 0, 5, 0, 0}
	c := condOf(t, "P31=Q5")
	if id, ok, err := checkRecord(rec, c); err != nil || !ok || id != 5 {
		t.Fatal(id, ok, err)
	}
	if _, ok, err := checkRecord(rec, condOf(t, "P31=Q6")); ok || err != nil {
		t.Fatal(ok, err)
	}
	if _, _, err := checkRecord(rec[:3], c); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
	if _, _, err := checkRecord(rec, condOf(t, "QUALIFIER:1=P31")); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
	short := []byte{0x11, 0x0a, 0, 0x88, 0, 5}
	if _, _, err := checkRecord(short, c); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
	badID := append([]byte{}, rec...)
	badID[2] = 0x00
	badID[3] = 0xa0
	if _, _, err := checkRecord(badID, c); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
}
