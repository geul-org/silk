//ff:func feature=sig type=matcher control=sequence
//ff:what TestPatternMatch — every word must agree under the mask
package sig

import "testing"

func TestPatternMatch(t *testing.T) {
	p := Pattern{Mask: []uint64{0xF, 0x1}, Want: []uint64{0x5, 0x1}}
	if !p.match([]uint64{0xF5, 0x3}) || p.match([]uint64{0x5, 0x2}) || p.match([]uint64{0x4, 0x1}) {
		t.Fatal("match")
	}
}
