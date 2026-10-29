//ff:func feature=cli type=helper control=sequence
//ff:what TestNewIndexFlags — every index flag parses into indexFlags, defaults included
package cli

import (
	"io"
	"testing"

	"github.com/park-jun-woo/silk/pkg/index"
)

func TestNewIndexFlags(t *testing.T) {
	var f indexFlags
	fs := newIndexFlags("x", io.Discard, &f)
	if fs.Parse([]string{"-index", "d", "-spec", "s", "-compact", "-full", "-idle", "-dry"}) != nil || f.dir != "d" || f.spec != "s" || !f.compact || !f.full || !f.idle || !f.dry || f.threads != index.DefaultThreads {
		t.Fatal(f)
	}
}
