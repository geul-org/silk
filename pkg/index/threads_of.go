//ff:func feature=index type=helper control=sequence
//ff:what threadsOf — worker count: the option, at least one
package index

func threadsOf(o Options) int {
	return max(o.Threads, 1)
}
