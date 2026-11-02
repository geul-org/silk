//ff:func feature=archive type=builder control=sequence
//ff:what testArchive — a packed fixture folder with two documents, opened
package archive

import (
	"context"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
)

func testArchive(t *testing.T) (string, *Archive) {
	t.Helper()
	s := fixture.Tok("Q1")
	data := fixture.Pack(t, []fixture.Doc{{ID: 0x5100000000000001, Version: 1, Body: []core.Record{fixture.Label(1, s, "en", "A")}}, {ID: 0x5100000000000002, Version: 1, Body: []core.Record{fixture.Label(1, s, "en", "B")}}}, nil)
	a, err := Open(context.Background(), data)
	if err != nil {
		t.Fatal(err)
	}
	return data, a
}
