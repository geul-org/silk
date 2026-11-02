//ff:func feature=grammar type=reader control=sequence
//ff:what testGrammar — the published grammar copy in testdata
package grammar

import (
	"context"
	"path/filepath"
	"runtime"
	"testing"
)

func testGrammar(t *testing.T) *Grammar {
	t.Helper()
	_, f, _, _ := runtime.Caller(0)
	g, err := Load(context.Background(), filepath.Join(filepath.Dir(f), "..", "..", "testdata", "grammar", "0x0000"))
	if err != nil {
		t.Fatal(err)
	}
	return g
}
