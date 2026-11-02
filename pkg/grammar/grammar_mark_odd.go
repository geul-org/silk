//ff:func feature=grammar type=builder control=iteration dimension=1
//ff:what Grammar.markOdd — byte-string pairs: two targets of one name that differ in the low bit; the higher is odd
package grammar

import "sort"

func (g *Grammar) markOdd() {
	for n, ts := range g.Names {
		sort.Slice(ts, func(i, j int) bool { return ts[i] < ts[j] })
		g.Names[n] = ts
		if len(ts) == 2 && ts[0]^ts[1] == 1 {
			g.Odd[ts[1]] = true
		}
	}
}
