//ff:type feature=idle type=model
//ff:what Sample — one reading of the system: 1-minute load average, CPU count, total and available memory in bytes
package idle

// Sample is what Read returns.
type Sample struct {
	Load1    float64
	CPUs     int
	MemTotal uint64
	MemAvail uint64
}
