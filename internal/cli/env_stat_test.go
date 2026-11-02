//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvStat — counts print; a missing index fails
package cli

import (
	"errors"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestEnvStat(t *testing.T) {
	data, f, _ := built(t)
	e, o := testEnv("")
	if err := e.stat(f, data); err != nil || !strings.Contains(o.String(), "generation") {
		t.Fatal(err, o)
	}
	f.index = t.TempDir()
	if err := e.stat(f, data); !errors.Is(err, index.ErrNoIndex) {
		t.Fatal(err)
	}
}
