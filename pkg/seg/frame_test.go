//ff:func feature=seg type=scanner control=sequence
//ff:what TestFrame — sizes that do not add up fail
package seg

import (
	"errors"
	"testing"
)

func TestFrame(t *testing.T) {
	fb := encodeFooter(testInfo(), 1, 1, make([]byte, 32))
	if _, err := frame("x", append(make([]byte, 6), fb...)); !errors.Is(err, ErrFormat) {
		t.Fatal(err)
	}
	if f, err := frame("x", append(make([]byte, 6+18), fb...)); err != nil || f.Rows != 1 {
		t.Fatal(err)
	}
}
