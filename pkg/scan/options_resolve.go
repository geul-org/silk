//ff:func feature=scan type=helper control=sequence
//ff:what Options.resolve — threads and window with defaults; never more threads than offsets, at least one
package scan

// resolve fills zero values and caps threads at n (minimum 1).
func (o Options) resolve(n int) (int, int) {
	th, win := o.Threads, o.Window
	if th <= 0 {
		th = DefaultThreads
	}
	if win <= 0 {
		win = DefaultWindow
	}
	return max(1, min(th, n)), win
}
