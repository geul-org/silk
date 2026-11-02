//ff:func feature=cli type=scanner control=sequence
//ff:what TestFlagsParseArgs — words and flags in any order; a bad flag fails
package cli

import (
	"io"
	"testing"
)

func TestFlagsParseArgs(t *testing.T) {
	f := newFlags("t", io.Discard)
	if p, err := f.parseArgs([]string{"a", "-threads", "2", "b"}); err != nil || len(p) != 2 || f.threads != 2 {
		t.Fatal(p, err)
	}
	if _, err := f.parseArgs([]string{"-nope"}); err == nil {
		t.Fatal("bad flag")
	}
}
