//ff:func feature=idle type=rule control=sequence
//ff:what Rule.Check — both shares are set, between 0 and 1
package idle

func (r Rule) Check() error {
	if r.CPUFree <= 0 || r.MemFree <= 0 || r.CPUFree >= 1 || r.MemFree >= 1 {
		return ErrRule
	}
	return nil
}
