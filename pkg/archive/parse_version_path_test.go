//ff:func feature=archive type=scanner control=iteration dimension=1
//ff:what TestParseVersionPath — version documents parse; anything else is false
package archive

import "testing"

func TestParseVersionPath(t *testing.T) {
	p := VersionPath("0x00010000", 0x51, 2)
	s, ok := parseVersionPath(p)
	if !ok || s.Issuer != "0x00010000" || s.ID != 0x51 || s.Version != 2 || p != "0x00010000/0x0000000000000051/0x00000002.geul" {
		t.Fatal(s, ok, p)
	}
	for _, x := range []string{"a/b", "0x1/x/y.geul", "0x1/0x0000000000000051/versions.geul", "0x1/0x0000000000000051/0x00000000.geul", "h/0xabc.geul", "0x1/0x0000000000000051/0x00000001.txt", ""} {
		if _, ok := parseVersionPath(x); ok {
			t.Fatal(x)
		}
	}
	for _, x := range []string{"0xABCD", "0x12", "abcd", "0xzzzz"} {
		if _, ok := parseHexWidth(x, 4); ok {
			t.Fatal(x)
		}
	}
}
