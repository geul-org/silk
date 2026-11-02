//ff:func feature=seg type=scanner control=sequence
//ff:what TestFileClose — a file closes once
package seg

import "testing"

func TestFileClose(t *testing.T) {
	p := writeSeg(t, "a").Path
	f, _ := Open(p)
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
}
