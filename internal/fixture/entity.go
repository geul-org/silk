//ff:func feature=fixture type=builder control=sequence
//ff:what entity — one random document
package fixture

import (
	"math/rand"

	"github.com/geul-org/gogeul/pkg/core"
)

func entity(r *rand.Rand, id uint64, subj []byte) Doc {
	d := Doc{ID: id, Version: 1}
	n := uint64(1)
	add := func(rec core.Record) {
		d.Body = append(d.Body, rec)
		n++
	}
	statements(r, subj, "P31", []string{"Q5", "Q5", "Q5", "Q16521", "Q515", "Q99"}, r.Intn(3), add, &n)
	statements(r, subj, "P106", []string{"Q82955", "Q36180", "Q33999", "Q169470", "Q98"}, r.Intn(4), add, &n)
	statements(r, subj, "P21", []string{"Q6581097", "Q6581072"}, r.Intn(2), add, &n)
	add(Triple(n, Tok("P569"), subj, dateOf(r), 0))
	add(Label(n, subj, []string{"en", "ko", "de"}[r.Intn(3)], names[r.Intn(len(names))]))
	if r.Intn(2) == 0 {
		add(Alias(n, subj, "en", names[r.Intn(len(names))]))
	}
	return d
}
