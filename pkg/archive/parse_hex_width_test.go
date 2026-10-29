//ff:func feature=archive type=decoder control=iteration dimension=1
//ff:what TestParseHexWidth — exact-width lowercase 0x hex parses; missing prefix, width, case or digits fail
package archive

import "testing"

func TestParseHexWidth(t *testing.T) {
	if v, ok := parseHexWidth("0x00ff", 4); !ok || v != 255 {
		t.Fatal(v, ok)
	}
	for _, s := range []string{"00ff", "0xff", "0x00FF", "0x00zz"} {
		if _, ok := parseHexWidth(s, 4); ok {
			t.Fatal(s)
		}
	}
}
