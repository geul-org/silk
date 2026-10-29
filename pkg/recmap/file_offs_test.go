//ff:func feature=recmap type=reader control=sequence
//ff:what TestFileOffs — a whole off[] reads back; a read failure is ErrFormat
package recmap

import (
	"bytes"
	"errors"
	"slices"
	"testing"
)

func TestFileOffs(t *testing.T) {
	m, _, _ := sampleMap(t)
	b, _ := m.Encode()
	f, _ := Read(bytes.NewReader(b), int64(len(b)))
	if got, err := f.offs(0, f.Header.Records); err != nil || !slices.Equal(got, m.Offs) {
		t.Fatal(err)
	}
	f.r = failReader{b: b, from: 0}
	if _, err := f.offs(0, 1); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
