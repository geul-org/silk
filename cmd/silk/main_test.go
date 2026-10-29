//ff:func feature=cli type=command control=sequence
//ff:what TestMainEntry — main passes the command line to cli.Run and exits with its code
package main

import (
	"os"
	"testing"
)

func TestMainEntry(t *testing.T) {
	code := -1
	exit = func(c int) { code = c }
	defer func() { exit = os.Exit }()
	args := os.Args
	defer func() { os.Args = args }()
	os.Args = []string{"silk", "nope"}
	main()
	if code != 2 {
		t.Fatal(code)
	}
}
