//go:build !unix

//ff:func feature=index type=helper control=sequence
//ff:what pidAlive — without a portable process check, any positive process ID counts as alive (locks and pins are then removed by hand)
package index

// pidAlive is conservative off Unix.
func pidAlive(pid int) bool {
	return pid > 0
}
