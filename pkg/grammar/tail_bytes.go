//ff:func feature=grammar type=helper control=iteration dimension=1
//ff:what tailBytes — bytes of the fixed raw fields after a rest field
package grammar

func tailBytes(fs []Field) int {
	n, rest := 0, false
	for _, f := range fs {
		if rest {
			n += 2 * f.Width
		}
		rest = rest || f.Width == RestWidth
	}
	return n
}
