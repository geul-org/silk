//ff:func feature=cli type=helper control=iteration dimension=1
//ff:what TestPrintReport — the summary with write amplification, merges, unchanged and paused notes; '-' without new bytes
package cli

import (
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestPrintReport(t *testing.T) {
	var b strings.Builder
	printReport(&b, "update", index.Report{NewBytes: 4, Written: 2, Tasks: []string{"TRIPLE flush"}, Unchanged: true, Paused: true})
	s := b.String()
	for _, w := range []string{"write amplification 0.500", "merge TRIPLE flush", "already matches", "paused"} {
		if !strings.Contains(s, w) {
			t.Fatal(w, s)
		}
	}
	b.Reset()
	printReport(&b, "compact", index.Report{})
	if !strings.Contains(b.String(), "write amplification -") {
		t.Fatal(b.String())
	}
}
