//ff:func feature=decl type=builder control=sequence
//ff:what rowsOfKind — one kind's extractor run on the test document
package decl

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/seg"
)

func rowsOfKind(t *testing.T, d Decl) []*seg.Buffer {
	t.Helper()
	x, err := NewExtractor(d, fixture.Grammar(t))
	if err != nil {
		t.Fatal(err)
	}
	b := buffers(t, d)
	x.Rows(testDoc(t), 9, b)
	return b
}
