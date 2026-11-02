//ff:func feature=archive type=reader control=iteration dimension=1
//ff:what TestSetStepsPaths — versions.geul entries of the grammar set are not version steps
package archive

import (
	"context"
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestSetStepsPaths(t *testing.T) {
	data := t.TempDir()
	wk := filepath.Join(data, ".well-known", "geul", "0x0000")
	if err := os.CopyFS(wk, os.DirFS(fixture.GrammarDir())); err != nil {
		t.Fatal(err)
	}
	a, err := Open(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	s := a.Sets[0]
	n := 0
	for _, d := range s.Data() {
		err := s.Steps(context.Background(), d.Obj, func(st Step) error { n++; return nil })
		if err != nil {
			t.Fatal(err)
		}
	}
	if n == 0 || n >= 430 {
		t.Fatal(n)
	}
}
