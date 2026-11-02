//ff:func feature=query type=helper control=sequence
//ff:what changed — the corpus after edits: document 0 rewritten with a job it did not have, document 1 removed, one document added
package query

import (
	"slices"

	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/internal/fixture"
)

func changed(docs []fixture.Doc) []fixture.Doc {
	out := slices.Clone(docs[2:])
	s := fixture.Tok("Q1000")
	out = append(out, fixture.Doc{ID: docs[0].ID, Version: 1, Body: []core.Record{
		fixture.Triple(1, fixture.Tok("P31"), s, fixture.Tok("Q5"), 0),
		fixture.Triple(2, fixture.Tok("P106"), s, fixture.Tok("Q33999"), 0),
		fixture.Label(3, s, "en", "Rewritten"),
	}})
	n := fixture.Tok("Q9999")
	return append(out, fixture.Doc{ID: 0x5100000000009999, Version: 1, Body: []core.Record{
		fixture.Triple(1, fixture.Tok("P31"), n, fixture.Tok("Q515"), 0),
		fixture.Label(2, n, "ko", "새 문서"),
	}})
}
