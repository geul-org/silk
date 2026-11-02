//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestParserFieldLine — common and block fields; every refusal
package sig

import (
	"strings"
	"testing"
)

func TestParserFieldLine(t *testing.T) {
	p := newParser()
	line := func(s string) error { return p.fieldLine(strings.Fields(s)) }
	if err := line("field a 0 4 TRIPLE P1 enum"); err != nil || len(p.d.Common) != 1 || p.field.Off != 0 {
		t.Fatal(err)
	}
	p.when = &When{}
	if err := line("field b 4 4 TRIPLE P2 enum"); err != nil || len(p.when.Fields) != 0 && p.when.Fields[0].Name != "b" {
		t.Fatal(err)
	}
	for _, s := range []string{"field a 0", "field c 0 4 LABEL P1 enum", "field c x 4 TRIPLE P1 enum", "field c 0 x TRIPLE P1 enum", "field c -1 4 TRIPLE P1 enum", "field c 0 0 TRIPLE P1 enum", "field c 0 65 TRIPLE P1 enum", "field a 0 4 TRIPLE P1 enum", "field c 0 4 TRIPLE X1 enum", "field c 0 4 TRIPLE P1 nope"} {
		if err := line(s); err == nil {
			t.Fatal(s)
		}
	}
}
