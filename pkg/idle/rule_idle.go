//ff:func feature=idle type=rule control=sequence
//ff:what Rule.Idle — (load + own threads) / CPUs leaves CPUFree unused and available memory is at least MemFree of the total
package idle

// Idle applies r to s.
func (r Rule) Idle(s Sample) bool {
	cpu, mem, th := r.CPUFree, r.MemFree, r.Threads
	if cpu <= 0 {
		cpu = DefaultFree
	}
	if mem <= 0 {
		mem = DefaultFree
	}
	if th <= 0 {
		th = 1
	}
	if s.CPUs <= 0 || s.MemTotal == 0 {
		return false
	}
	busy := (s.Load1 + float64(th)) / float64(s.CPUs)
	return busy <= 1-cpu && float64(s.MemAvail) >= mem*float64(s.MemTotal)
}
