//ff:func feature=idle type=decoder control=sequence
//ff:what TestParseMeminfo — MemTotal and MemAvailable in bytes; a bad number or no MemTotal fails
package idle

import "testing"

func TestParseMeminfo(t *testing.T) {
	tot, av, err := parseMeminfo([]byte("MemTotal:  16 kB\nMemFree: 1 kB\nMemAvailable:   8 kB\n\n"))
	if err != nil || tot != 16384 || av != 8192 {
		t.Fatal(tot, av, err)
	}
	if _, _, err := parseMeminfo([]byte("MemTotal: x kB\n")); err == nil {
		t.Fatal("bad number")
	}
	if _, _, err := parseMeminfo([]byte("MemAvailable: 1 kB\n")); err == nil {
		t.Fatal("no MemTotal")
	}
}
