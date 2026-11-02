//ff:func feature=query type=builder control=iteration dimension=1
//ff:what extend — every prefix with every value's slot; a prefix predicate adds its bytes unpadded
package query

func extend(ps [][]byte, q Pred, rest bool, w int) [][]byte {
	var out [][]byte
	for _, base := range ps {
		for _, v := range q.Vals {
			out = append(out, appendKey(base, q, rest, v, w))
		}
	}
	return out
}
