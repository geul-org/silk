//ff:func feature=recmap type=reader control=sequence
//ff:what TestFileOffsets — a key's offsets come from its slice of off[]; an absent key has none
package recmap

import (
	"bytes"
	"slices"
	"testing"
)

func TestFileOffsets(t *testing.T) {
	m, _, _ := sampleMap(t)
	b, _ := m.Encode()
	f, _ := Read(bytes.NewReader(b), int64(len(b)))
	k := m.Keys[len(m.Keys)-1]
	if got, err := f.Offsets(k.Key); err != nil || !slices.Equal(got, m.Offs[k.Start:k.Start+k.Count]) {
		t.Fatal(got, err)
	}
	if got, err := f.Offsets(keySpace - 1); err != nil || got != nil {
		t.Fatal(got, err)
	}
}
