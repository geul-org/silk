//ff:func feature=archive type=decoder control=iteration dimension=1
//ff:what TestParseVersionPath — canonical version paths parse; versions.geul, h/, named, short IDs and version 0 do not
package archive

import "testing"

func TestParseVersionPath(t *testing.T) {
	d, ok := parseVersionPath("0x00010000/0x5100000000001d52/0x0000000a.geul")
	if !ok || d != (Doc{Issuer: "0x00010000", ID: 0x5100000000001d52, Version: 10}) {
		t.Fatal(d, ok)
	}
	for _, p := range []string{
		"0x00010000/0x5100000000001d52/versions.geul",
		"h/0x00.geul",
		"0x00010000/documents.geul",
		"x0001/0x5100000000001d52/0x00000001.geul",
		"0x00010000/0x5100000000001d52/0x00000001.txt",
		"0x00010000/0x51/0x00000001.geul",
		"0x00010000/0x5100000000001d52/0x00000000.geul",
	} {
		if _, ok := parseVersionPath(p); ok {
			t.Fatal(p)
		}
	}
}
