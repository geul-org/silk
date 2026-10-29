//ff:func feature=index type=writer control=sequence
//ff:what TestPinGen — a pin file named for the generation and this process; a blocked pin directory fails
package index

import (
	"os"
	"strings"
	"testing"
)

func TestPinGen(t *testing.T) {
	d := t.TempDir()
	p, err := pinGen(d, 7)
	if _, serr := os.Stat(p); err != nil || serr != nil || !strings.Contains(p, "/pin/7-") {
		t.Fatal(p, err)
	}
	e := t.TempDir()
	block(e, dirPin)
	if _, err := pinGen(e, 1); err == nil {
		t.Fatal("blocked pin directory")
	}
}
