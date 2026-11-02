//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvWrite — build and update report; a missing limit, archive or broken index fails
package cli

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestEnvWrite(t *testing.T) {
	data, f, _ := built(t)
	e, o := testEnv("")
	if err := e.write(f, data, false); err != nil || !strings.Contains(o.String(), "new documents 0") {
		t.Fatal(err, o)
	}
	if err := e.write(f, t.TempDir(), true); err == nil {
		t.Fatal("no archive")
	}
	file := filepath.Join(t.TempDir(), "f")
	os.WriteFile(file, nil, 0o644)
	f.index = filepath.Join(file, "x")
	if err := e.write(f, data, true); err == nil {
		t.Fatal("under a file")
	}
	f.mem = ""
	if err := e.write(f, data, true); !errors.Is(err, index.ErrLimit) {
		t.Fatal(err)
	}
}
