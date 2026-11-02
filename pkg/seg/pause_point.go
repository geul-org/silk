//ff:func feature=seg type=helper control=sequence
//ff:what pausePoint — every 2^20 rows, give the caller its pause
package seg

func pausePoint(hk Hooks, n int) error {
	if hk.Pause == nil || n&(1<<20-1) != 0 {
		return nil
	}
	return hk.Pause()
}
