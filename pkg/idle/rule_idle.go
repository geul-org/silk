//ff:func feature=idle type=rule control=sequence
//ff:what Rule.Idle — (load + own threads) / CPUs leaves CPUFree unused and available memory is at least MemFree of the total
package idle

func (r Rule) Idle(s Sample) bool {
	if s.CPUs <= 0 || s.MemTotal == 0 {
		return false
	}
	busy := (s.Load1 + float64(max(r.Threads, 1))) / float64(s.CPUs)
	return busy <= 1-r.CPUFree && float64(s.MemAvail) >= r.MemFree*float64(s.MemTotal)
}
