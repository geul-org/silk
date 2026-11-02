//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what project — a packed folder with one document, a grammar flag and definition files; created sections in -index
package cli

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
)

func project(t *testing.T) (string, []string) {
	t.Helper()
	s := fixture.Tok("Q1")
	data := fixture.Pack(t, []fixture.Doc{{ID: 1, Version: 1, Body: []core.Record{fixture.Triple(1, fixture.Tok("P31"), s, fixture.Tok("Q5"), 0), fixture.Label(2, s, "en", "Albert")}}}, nil)
	dir := t.TempDir()
	def, sem := filepath.Join(dir, "d"), filepath.Join(dir, "s")
	os.WriteFile(def, []byte(defText), 0o644)
	os.WriteFile(sem, []byte(semText), 0o644)
	ix := []string{"-index", filepath.Join(dir, "ix")}
	g := append([]string{"-grammar", fixture.GrammarDir()}, ix...)
	for _, a := range [][]string{
		{"create", "index", "triple", "TRIPLE", "-part", "predicate", "-key", "object", "-width", "6"},
		{"create", "text", "label", "LABEL"},
		{"create", "silk", "entity", "-def", def},
		{"create", "composite", "sem", "LABEL", "-def", sem},
	} {
		if c, _, e := silkg("", append(append(a, g...), data)...); c != 0 {
			t.Fatal(a, e)
		}
	}
	return data, ix
}
