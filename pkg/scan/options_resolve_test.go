//ff:func feature=scan type=helper control=sequence
//ff:what TestOptionsResolve — defaults fill zeros; threads capped by work, at least one
package scan

import "testing"

func TestOptionsResolve(t *testing.T) {
	if th, w := (Options{}).resolve(100); th != DefaultThreads || w != DefaultWindow {
		t.Fatal(th, w)
	}
	if th, w := (Options{Threads: 4, Window: 9}).resolve(2); th != 2 || w != 9 {
		t.Fatal(th, w)
	}
	if th, _ := (Options{}).resolve(0); th != 1 {
		t.Fatal(th)
	}
}
