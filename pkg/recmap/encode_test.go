//ff:func feature=recmap type=encoder control=sequence
//ff:what TestMapEncode — counts and body hash are filled and the bytes read back to the same map; a bad Issuer fails
package recmap

import (
	"bytes"
	"crypto/sha256"
	"slices"
	"testing"
)

func TestMapEncode(t *testing.T) {
	m, _, _ := sampleMap(t)
	m.Header.Records = 0
	b, err := m.Encode()
	if err != nil || m.Header.Records != uint64(len(m.Offs)) || m.Header.Sum != sha256.Sum256(b[headerSize:]) {
		t.Fatal(err)
	}
	f, err := Read(bytes.NewReader(b), int64(len(b)))
	if err != nil || f.Header != m.Header || !slices.Equal(f.Keys, m.Keys) || !slices.Equal(f.Docs, m.Docs) {
		t.Fatal(f, err)
	}
	m.Header.Issuer = "bad"
	if _, err := m.Encode(); err == nil {
		t.Fatal("bad Issuer must fail")
	}
}
