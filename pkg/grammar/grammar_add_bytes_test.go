//ff:func feature=grammar type=decoder control=sequence
//ff:what TestGrammarAddBytes — a document's body adds; bytes that are not one document fail
package grammar

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestGrammarAddBytes(t *testing.T) {
	g := &Grammar{Fields: map[uint16][]Field{}, Names: map[string][]uint16{}, Langs: map[string]uint16{}}
	b, err := core.EncodeDocument(nil, nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := g.addBytes(b); err != nil {
		t.Fatal(err)
	}
	if err := g.addBytes(b[:len(b)-2]); err == nil {
		t.Fatal("cut document")
	}
}
