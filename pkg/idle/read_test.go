//ff:func feature=idle type=reader control=sequence
//ff:what TestRead — on Linux a Sample with CPUs and memory, elsewhere ErrUnsupported
package idle

import (
	"errors"
	"testing"
)

func TestRead(t *testing.T) {
	s, err := Read()
	if err != nil && !errors.Is(err, ErrUnsupported) {
		t.Fatal(err)
	}
	if err == nil && (s.CPUs < 1 || s.MemTotal == 0) {
		t.Fatal(s)
	}
}
