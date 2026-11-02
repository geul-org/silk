//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvCompact — options, the idle rule and a broken index stop it; merges are reported
package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/idle"
	"github.com/geul-org/silk/pkg/index"
)

func TestEnvCompact(t *testing.T) {
	data, f, _ := built(t)
	e, o := testEnv("")
	f.full = true
	if err := e.compact(f, data); err != nil || !strings.Contains(o.String(), "merges") {
		t.Fatal(err, o)
	}
	f.idleCPU = 2
	if err := e.compact(f, data); !errors.Is(err, idle.ErrRule) {
		t.Fatal(err)
	}
	f.idleCPU, f.mem = 0, ""
	if err := e.compact(f, data); !errors.Is(err, index.ErrLimit) {
		t.Fatal(err)
	}
	f.mem, f.index = "64M", filepath.Join(data, ".well-known", "geul", "latest", "x")
	file := filepath.Join(t.TempDir(), "f")
	os.WriteFile(file, nil, 0o644)
	f.index = filepath.Join(file, "x")
	if err := e.compact(f, data); err == nil {
		t.Fatal("under a file")
	}
}
