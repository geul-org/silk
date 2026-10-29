//ff:func feature=spec type=decoder control=sequence
//ff:what TestSubRecord — the language record of a UTF8-LANGUAGE-TEXT (even and odd), false for another record or a broken payload
package spec

import (
	"bytes"
	"testing"
)

func TestSubRecord(t *testing.T) {
	if f, ok := subRecord(hexBytes("0c4a 08c8 0783 446f 6e61 6c64 2054 7275 6d70")); !ok || !bytes.Equal(f, hexBytes("08c8 0783")) {
		t.Fatal(f)
	}
	if f, ok := subRecord(hexBytes("0e4b 0947 6173 7400 5300")); !ok || len(f) != 6 {
		t.Fatal(f)
	}
	if _, ok := subRecord(hexBytes("0088 0001")); ok {
		t.Fatal("LOCAL Reference")
	}
	if _, ok := subRecord(hexBytes("084a 0f00")); ok {
		t.Fatal("broken payload")
	}
}
