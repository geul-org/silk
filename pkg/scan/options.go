//ff:type feature=scan type=model
//ff:what Options — scan threads and window size (0 takes DefaultThreads, DefaultWindow)
package scan

// Options tune Records.
type Options struct {
	Threads int
	Window  int
}
