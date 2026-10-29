//ff:func feature=index type=rule control=sequence
//ff:what TestCheckDoc — a real document passes; an odd byte count or bytes that are not a document fail
package index

import (
	"errors"
	"testing"
)

func TestCheckDoc(t *testing.T) {
	g, _, _, docs := testIngest(t, 1)
	b, _ := g.readDoc(docs[0])
	if err := checkDoc(b); err != nil {
		t.Fatal(err)
	}
	if err := checkDoc(b[:3]); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if err := checkDoc(b[2:]); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
