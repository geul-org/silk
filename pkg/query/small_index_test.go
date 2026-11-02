//ff:func feature=query type=helper control=sequence
//ff:what smallIndex — a small document set exercising every section kind
package query

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
	"github.com/geul-org/silk/pkg/archive"
	"github.com/geul-org/silk/pkg/index"
)

func smallIndex(t *testing.T) (*archive.Archive, string, *index.Index) {
	t.Helper()
	s := fixture.Tok("Q1")
	body := []core.Record{
		fixture.Triple(1, fixture.Tok("P31"), s, fixture.Tok("Q5"), 0),
		fixture.Triple(2, fixture.Tok("P106"), s, fixture.Tok("Q169470"), 0),
		fixture.Triple(3, fixture.Tok("P21"), s, []byte{0, 1}, 0),
		fixture.Label(4, s, "en", "Albert Einstein"),
		fixture.LabelShort(5, s, 0x0784, "Short"),
		fixture.Alias(6, s, "en", "Einstein"),
		fixture.Bind(7, s, "albert"),
	}
	body = append(body, fixture.LabelLocal(8, 50, s, "de", "Lokal")...)
	o := fixture.Tok("Q2")
	other := []core.Record{fixture.Triple(1, fixture.Tok("P31"), o, fixture.Tok("Q99"), 0), fixture.Triple(2, fixture.Tok("P106"), o, fixture.Tok("Q98"), 0), fixture.Label(3, o, "en", "Albert Other")}
	return openIndex(t, fixture.Pack(t, []fixture.Doc{{ID: 1, Version: 1, Body: body}, {ID: 2, Version: 1, Body: other}}, nil))
}
