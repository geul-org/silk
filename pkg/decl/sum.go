//ff:func feature=decl type=helper control=iteration dimension=1
//ff:what sum — add ints
package decl

func sum(v []int) int {
	n := 0
	for _, x := range v {
		n += x
	}
	return n
}
