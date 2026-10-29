//ff:func feature=idle type=decoder control=sequence
//ff:what TestParseLoadavg — the first field of /proc/loadavg; empty input and a non-number fail
package idle

import "testing"

func TestParseLoadavg(t *testing.T) {
	if v, err := parseLoadavg([]byte("0.17 0.10 0.09 1/631 3321009\n")); err != nil || v != 0.17 {
		t.Fatal(v, err)
	}
	if _, err := parseLoadavg(nil); err == nil {
		t.Fatal("empty")
	}
	if _, err := parseLoadavg([]byte("x 1")); err == nil {
		t.Fatal("not a number")
	}
}
