//ff:func feature=fixture type=scanner control=sequence
//ff:what TestEntity — a document's body holds a birth and a label; statements add records and count them
package fixture

import (
	"math/rand"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEntity(t *testing.T) {
	r := rand.New(rand.NewSource(2))
	d := entity(r, 9, Tok("Q9"))
	if d.ID != 9 || len(d.Body) < 2 {
		t.Fatal(d)
	}
	var out []core.Record
	n := uint64(1)
	statements(r, Tok("Q1"), "P31", []string{"Q5"}, 40, func(c core.Record) { out = append(out, c); n++ }, &n)
	if len(out) != 40 || n != 41 {
		t.Fatal(len(out), n)
	}
}
