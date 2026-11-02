//ff:func feature=seg type=scanner control=sequence
//ff:what TestPausePoint — the hook runs at multiples of 2^20 rows only
package seg

import "testing"

func TestPausePoint(t *testing.T) {
	n := 0
	hk := Hooks{Pause: func() error { n++; return nil }}
	pausePoint(hk, 1<<20)
	pausePoint(hk, 3)
	if pausePoint(Hooks{}, 0) != nil || n != 1 {
		t.Fatal(n)
	}
}
