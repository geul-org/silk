//ff:func feature=cli type=scanner control=iteration dimension=1
//ff:what TestUsageErrors — wrong argument counts and unreadable inputs
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestUsageErrors(t *testing.T) {
	data, ix := project(t)
	dir := t.TempDir()
	cases := [][]string{
		{"create"},
		{"create", "-nope"},
		{"create", "silk", "x", "-def", filepath.Join(dir, "none"), "-grammar", fixture.GrammarDir(), data},
		{"create", "composite", "x", "LABEL", "-def", filepath.Join(dir, "none"), "-grammar", fixture.GrammarDir(), data},
		{"create", "index", "x", "-grammar", fixture.GrammarDir(), data},
		{"create", "nope", "x", "LABEL", "-grammar", fixture.GrammarDir(), data},
		{"create", "index", "x", "LABEL", "-width", "z", "-grammar", fixture.GrammarDir(), data},
		{"create", "text", "x", "LABEL", dir},
		{"create", "text", "x", "LABEL", "-grammar", dir, data},
		{"create", "text", "x", "LABEL", data},
		{"drop", "-nope"},
		{"drop", "a"},
		{"list", "-nope"},
		{"list"},
		{"list", "-index", dir, data},
		{"index"},
		{"index", "stat", "-nope", data},
		{"index", "stat"},
		{"index", "stat", "-index", dir, data},
		{"index", "build", "-mem", "x", data},
		{"index", "build", "-mem", "1K", dir},
		{"index", "update", "-mem", "1K", "-index", filepath.Join(dir, "x", "y"), filepath.Join(dir, "none")},
		{"index", "verify", "-mem", "1K", dir},
		{"index", "verify", "-mem", "1K", "-index", dir, data},
		{"index", "compact", "-mem", "x", data},
		{"index", "compact", "-mem", "1K", "-index", filepath.Join(dir, "f"), data},
		{"query"},
		{"query", "-nope"},
		{"query", "TRIPLE", dir},
		{"query", "TRIPLE", "-index", dir, data},
	}
	os.WriteFile(filepath.Join(dir, "f"), nil, 0o644)
	for _, a := range cases {
		if c, _, _ := silkg("", a...); c == 0 {
			t.Fatal(a)
		}
	}
	if c, _, e := silkg("", append(append([]string{"index", "build", "-mem", "64M"}, ix...), data)...); c != 0 {
		t.Fatal(e)
	}
	if c, _, e := silkg("", "index", "build", "-mem", "64M", "-index", filepath.Join(dir, "other"), data); c != 0 {
		t.Fatal("empty index builds nothing", e)
	}
}
