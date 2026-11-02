//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvVerify — a sound index is ok; a missing limit, archive or index fails
package cli

import (
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestEnvVerify(t *testing.T) {
	data, f, _ := built(t)
	e, o := testEnv("")
	if err := e.verify(f, data); err != nil || !strings.Contains(o.String(), "ok") {
		t.Fatal(err, o)
	}
	if err := e.verify(f, t.TempDir()); err == nil {
		t.Fatal("no archive")
	}
	f.index = t.TempDir()
	if err := e.verify(f, data); err == nil {
		t.Fatal("no index")
	}
	f.mem = ""
	if err := e.verify(f, data); !errors.Is(err, index.ErrLimit) {
		t.Fatal(err)
	}
}
