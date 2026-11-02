//ff:func feature=query type=scanner control=sequence
//ff:what TestPlanKeysOK — a covered index, a text section, an uncovered index
package query

import (
	"testing"

	"github.com/geul-org/silk/pkg/decl"
)

func TestPlanKeysOK(t *testing.T) {
	if !(&Plan{Covered: true}).keysOK() || !(&Plan{Decl: decl.Decl{Kind: decl.KindText}}).keysOK() || (&Plan{Decl: decl.Decl{Kind: decl.KindIndex}}).keysOK() {
		t.Fatal("keysOK")
	}
}
