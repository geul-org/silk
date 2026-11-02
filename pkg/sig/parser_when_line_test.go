//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestParserWhenLine — a common enum's code opens a block; other targets fail
package sig

import (
	"strings"
	"testing"
)

func TestParserWhenLine(t *testing.T) {
	p := newParser()
	p.fieldLine(strings.Fields("field a 0 4 TRIPLE P1 enum"))
	p.codeLine(strings.Fields("code c 4 Q1"))
	p.fieldLine(strings.Fields("field f 4 1 TRIPLE P2 flag"))
	if err := p.whenLine(strings.Fields("when a=c")); err != nil || p.when.Entry.Name != "c" {
		t.Fatal(err)
	}
	for _, s := range []string{"when", "when a", "when z=c", "when f=c", "when a=z"} {
		if err := p.whenLine(strings.Fields(s)); err == nil {
			t.Fatal(s)
		}
	}
}
