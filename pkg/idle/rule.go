//ff:type feature=idle type=model
//ff:what Rule — when work may run: CPU and memory shares to leave free, and the work's own thread count
package idle

type Rule struct {
	CPUFree float64
	MemFree float64
	Threads int
}
