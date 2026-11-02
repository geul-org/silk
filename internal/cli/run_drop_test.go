//ff:func feature=cli type=scanner control=iteration dimension=1
//ff:what TestRunDrop — a declared name goes; bad arguments fail
package cli

import "testing"

func TestRunDrop(t *testing.T) {
	e, _ := testEnv("")
	data, ix := project(t)
	if err := runDrop(e, append([]string{"sem"}, append(ix, data)...)); err != nil {
		t.Fatal(err)
	}
	for _, a := range [][]string{{"-nope"}, {"x"}} {
		if err := runDrop(e, a); err == nil {
			t.Fatal(a)
		}
	}
}
