//ff:func feature=index type=helper control=iteration dimension=1
//ff:what anyTrue — whether a mask has a true entry
package index

// anyTrue reports any(v).
func anyTrue(v []bool) bool {
	for _, b := range v {
		if b {
			return true
		}
	}
	return false
}
