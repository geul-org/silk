//ff:func feature=spec type=decoder control=sequence
//ff:what TestHeadLen — 2 bytes for a direct length, 4 and 6 with a Length extension, 0 without a TAG WORD
package spec

import "testing"

func TestHeadLen(t *testing.T) {
	if headLen(hexBytes("148a")) != 2 || headLen(hexBytes("170a 0026")) != 4 || headLen(hexBytes("178a 0001 0000")) != 6 || headLen([]byte{1}) != 0 {
		t.Fatal("headLen")
	}
}
