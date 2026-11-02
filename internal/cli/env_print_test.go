//ff:func feature=cli type=scanner control=sequence
//ff:what TestEnvPrint — lines carry the lead; a failing hit stops
package cli

import (
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/query"
)

func TestEnvPrint(t *testing.T) {
	_, _, ix := built(t)
	e, o := testEnv("")
	if err := e.print(ix, query.ModeDocs, "7\t", []query.Hit{{Doc: 0}}); err != nil || !strings.HasPrefix(o.String(), "7\t") {
		t.Fatal(err, o)
	}
	if err := e.print(ix, query.ModeRecords, "", []query.Hit{{Doc: 0, HasOff: true, Off: 1 << 30}}); err == nil {
		t.Fatal("bad hit")
	}
}
