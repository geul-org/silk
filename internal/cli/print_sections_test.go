//ff:func feature=cli type=scanner control=sequence
//ff:what TestPrintSections — one line per section, built or pending
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/geul-org/silk/pkg/index"
)

func TestPrintSections(t *testing.T) {
	var b bytes.Buffer
	printSections(&b, index.Stat{Sections: []index.SectionStat{{Name: "a", Pending: true}, {Name: "b"}}})
	if s := b.String(); !strings.Contains(s, "pending") || !strings.Contains(s, "built") {
		t.Fatal(s)
	}
}
