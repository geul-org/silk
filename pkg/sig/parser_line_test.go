//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestParserLine — dispatch by first word; blanks pass; unknown lines fail
package sig

import (
	"strings"
	"testing"
)

func TestParserLine(t *testing.T) {
	p := newParser()
	for _, l := range []string{"", "silk a 16", "field x 0 4 TRIPLE P1 enum", "code c 4 Q1", "when x=c"} {
		if err := p.line(strings.Fields(l)); err != nil {
			t.Fatal(l, err)
		}
	}
	if err := p.line([]string{"nope"}); err == nil {
		t.Fatal("unknown")
	}
}
