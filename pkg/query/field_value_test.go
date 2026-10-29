//ff:func feature=query type=decoder control=sequence
//ff:what TestFieldValue — a whole field, the language record of a language-tagged text; false for a sub-record of a Reference or of a broken payload
package query

import (
	"bytes"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFieldValue(t *testing.T) {
	ref := []core.Word{0x00a0, 0x589e}
	text := []core.Word{0x0c4a, 0x08c8, 0x0783, 0x446f, 0x6e61, 0x6c64, 0x2054, 0x7275, 0x6d70}
	fs := [][]core.Word{ref, text, {0x084a, 0x0f00}}
	if v, ok := fieldValue(fs, []int{0}); !ok || !bytes.Equal(v, []byte{0, 0xa0, 0x58, 0x9e}) {
		t.Fatal(v)
	}
	if v, ok := fieldValue(fs, []int{1, 0}); !ok || !bytes.Equal(v, []byte{0x08, 0xc8, 0x07, 0x83}) {
		t.Fatalf("%x", v)
	}
	if _, ok := fieldValue(fs, []int{0, 0}); ok {
		t.Fatal("Reference has no sub-record")
	}
	if _, ok := fieldValue([][]core.Word{{0x08ca, 0x0fff}}, []int{0, 0}); ok {
		t.Fatal("broken payload")
	}
}
