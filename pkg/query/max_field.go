//ff:func feature=query type=helper control=iteration dimension=1
//ff:what maxField — the highest field number a condition reads (0, the Edge Id, at least)
package query

// maxField tells checkRecord how many fields to decode.
func maxField(c Cond) int {
	n := 0
	for _, f := range c.Fields {
		n = max(n, f.Path[0])
	}
	return n
}
