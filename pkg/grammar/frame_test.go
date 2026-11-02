//ff:func feature=grammar type=scanner control=sequence
//ff:what TestFrame — records frame and a cut one fails
package grammar

import (
	"bytes"
	"errors"
	"slices"
	"testing"
)

func TestFrame(t *testing.T) {
	r := testLabel()
	d, err := Frame(append(slices.Clone(r), r...))
	if err != nil || len(d.Recs) != 2 || d.Recs[1].Off != len(r) || d.Recs[1].Tag != 0x1004 || !bytes.Equal(d.At(d.Recs[0]), r) {
		t.Fatal(d, err)
	}
	if _, err := Frame(r[:len(r)-2]); !errors.Is(err, ErrFrame) {
		t.Fatal(err)
	}
	if _, err := Frame([]byte{0xFF, 0}); !errors.Is(err, ErrFrame) {
		t.Fatal(err)
	}
}
