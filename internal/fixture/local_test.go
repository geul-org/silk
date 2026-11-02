//ff:func feature=fixture type=scanner control=sequence
//ff:what TestLocal — a LOCAL record
package fixture

import (
	"bytes"
	"testing"
)

func TestLocal(t *testing.T) {
	if !bytes.Equal(Local(3), []byte{0, 0x88, 0, 3}) {
		t.Fatal(Local(3))
	}
}
