//ff:func feature=cli type=helper control=sequence
//ff:what TestNewFlags — the set has -index and reports errors to stderr
package cli

import (
	"bytes"
	"testing"
)

func TestNewFlags(t *testing.T) {
	var errb bytes.Buffer
	fs, dir := newFlags("x", &errb)
	if fs.Parse([]string{"-index", "/i", "a"}) != nil || *dir != "/i" || fs.Arg(0) != "a" {
		t.Fatal(*dir)
	}
	if fs.Parse([]string{"-bad"}) == nil || errb.Len() == 0 {
		t.Fatal("bad flag")
	}
}
