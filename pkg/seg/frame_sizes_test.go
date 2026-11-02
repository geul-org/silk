//ff:func feature=seg type=scanner control=sequence
//ff:what TestFrameSizes — a row count too large for the file, a partition count too large, and a sum that is off
package seg

import (
	"errors"
	"testing"
)

func TestFrameSizes(t *testing.T) {
	big := encodeFooter(testInfo(), 0, 1<<40, make([]byte, 32))
	if _, err := frame("x", append(make([]byte, 6), big...)); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	many := encodeFooter(testInfo(), 1<<40, 0, make([]byte, 32))
	if _, err := frame("x", append(make([]byte, 6), many...)); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	bad := encodeFooter(Info{}, 0, 0, make([]byte, 32))
	if _, err := frame("x", bad); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
}
