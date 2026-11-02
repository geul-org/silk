//ff:func feature=fixture type=builder control=iteration dimension=1
//ff:what Corpus — n random entities with Subject Q1000+i: every statement kind a silk meets, by a seeded generator
package fixture

import (
	"math/rand"
	"strconv"
)

func Corpus(seed int64, n int) []Doc {
	r := rand.New(rand.NewSource(seed))
	out := make([]Doc, 0, n)
	for i := range n {
		out = append(out, entity(r, uint64(0x5100000000000000)+uint64(1000+i), Tok("Q"+strconv.Itoa(1000+i))))
	}
	return out
}
