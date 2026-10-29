//ff:func feature=seg type=reader control=sequence
//ff:what TestFileClose — closing works once and fails the second time
package seg

import "testing"

func TestFileClose(t *testing.T) {
	f := sampleSeg(t, true)
	if err := f.Close(); err != nil {
		t.Fatal(err)
	}
	if err := f.Close(); err == nil {
		t.Fatal("closed twice")
	}
}
