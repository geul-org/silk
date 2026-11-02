//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvLine — marks and Edge Ids per unit; a document that cannot be read surfaces
package cli

import (
	"context"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/query"
)

func TestEnvLine(t *testing.T) {
	_, _, ix := built(t)
	e, _ := testEnv("")
	s, err := e.line(ix, query.ModeKeys, query.Hit{Doc: 0, Key: []byte{1}, QID: 7, Loose: true, Unknown: true})
	if err != nil || !strings.Contains(s, " ~ ? 01 qid=00000007") {
		t.Fatal(s, err)
	}
	s, err = e.line(ix, query.ModeDocs, query.Hit{Doc: 0})
	if err != nil || strings.Contains(s, "#") {
		t.Fatal(s, err)
	}
	q, _ := query.Parse(ix, []string{"TRIPLE", "predicate=P31", "object=Q5"})
	r, _ := query.Run(context.Background(), ix, q)
	if s, err := e.line(ix, query.ModeDocs, r.Hits[0]); err != nil || !strings.Contains(s, " #1") {
		t.Fatal(s, err)
	}
	if _, err := e.line(ix, query.ModeRecords, query.Hit{Doc: 0, HasOff: true, Off: 1 << 30}); err == nil {
		t.Fatal("past the document")
	}
}
