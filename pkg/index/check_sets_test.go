//ff:func feature=index type=rule control=sequence
//ff:what TestCheckSets — the archive's own sets pass; another count or another generation is ErrStale
package index

import (
	"errors"
	"testing"
)

func TestCheckSets(t *testing.T) {
	a, _, _ := built(t)
	s := setRefs(a)
	if err := checkSets(a, s); err != nil {
		t.Fatal(err)
	}
	if err := checkSets(a, nil); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
	s[0].Generation++
	if err := checkSets(a, s); !errors.Is(err, ErrStale) {
		t.Fatal(err)
	}
}
