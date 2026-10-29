//ff:func feature=cli type=helper control=sequence
//ff:what TestNewQueryFlags — every query flag parses, repeatable conditions in order
package cli

import (
	"io"
	"testing"
)

func TestNewQueryFlags(t *testing.T) {
	var f queryFlags
	fs := newQueryFlags(io.Discard, &f)
	err := fs.Parse([]string{"-index", "d", "-threads", "2", "-count", "-explain", "-triple", "a", "-triple", "b", "-match", "m", "-label", "ko"})
	if err != nil || f.dir != "d" || f.threads != 2 || !f.count || !f.explain || len(f.triples) != 2 || f.matches[0] != "m" || f.labels[0] != "ko" {
		t.Fatal(err, f)
	}
}
