//ff:func feature=index type=decoder control=sequence
//ff:what TestFileNumber — the hex file counter of a name; 0 for anything else
package index

import "testing"

func TestFileNumber(t *testing.T) {
	if fileNumber("seg/000000000000002a.sseg") != 42 || fileNumber("seg/x.sseg") != 0 {
		t.Fatal("fileNumber")
	}
}
