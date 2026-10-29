//ff:func feature=index type=helper control=sequence
//ff:what TestAnyTrue — true when any entry is
package index

import "testing"

func TestAnyTrue(t *testing.T) {
	if !anyTrue([]bool{false, true}) || anyTrue([]bool{false}) || anyTrue(nil) {
		t.Fatal("anyTrue")
	}
}
