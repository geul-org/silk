//ff:func feature=fixture type=scanner control=iteration dimension=1
//ff:what TestDateOf — every date form, reproducibly
package fixture

import (
	"math/rand"
	"testing"

	"github.com/geul-org/silk/pkg/grammar"
)

func TestDateOf(t *testing.T) {
	r := rand.New(rand.NewSource(1))
	seen := map[uint16]bool{}
	for range 40 {
		seen[grammar.Target(dateOf(r))] = true
	}
	if len(seen) < 3 {
		t.Fatal(seen)
	}
}
