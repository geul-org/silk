//ff:func feature=fixture type=builder control=iteration dimension=1
//ff:what statements — count statements on a predicate with random rank; sometimes a marker
package fixture

import (
	"math/rand"

	"github.com/geul-org/gogeul/pkg/core"
)

func statements(r *rand.Rand, subj []byte, pred string, vals []string, count int, add func(core.Record), n *uint64) {
	for range count {
		obj := Tok(vals[r.Intn(len(vals))])
		switch r.Intn(8) {
		case 0:
			obj = []byte{0, 1}
		case 1:
			obj = []byte{0, 2}
		}
		add(Triple(*n, Tok(pred), subj, obj, []uint16{0, 0, 1, 2}[r.Intn(4)]))
	}
}
