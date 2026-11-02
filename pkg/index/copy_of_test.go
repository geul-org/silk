//ff:func feature=index type=helper control=sequence
//ff:what copyOf — an updater with the same fields and a fresh counter, for a test to bend
package index

func copyOf(u *updater) *updater {
	return &updater{ar: u.ar, dir: u.dir, o: u.o, m: u.m, tmp: u.tmp, live: u.live, first: u.first}
}
