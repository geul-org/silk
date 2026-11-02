//ff:func feature=sig type=scanner control=sequence
//ff:what TestParserCodeLine — a code goes to the open field; without one, or too short, it fails
package sig

import (
	"strings"
	"testing"
)

func TestParserCodeLine(t *testing.T) {
	p := newParser()
	if err := p.codeLine([]string{"code", "c", "4"}); err == nil {
		t.Fatal("no field")
	}
	p.fieldLine(strings.Fields("field a 0 4 TRIPLE P1 enum"))
	if err := p.codeLine([]string{"code", "c"}); err == nil {
		t.Fatal("short")
	}
	if err := p.codeLine([]string{"code", "c", "x", "Q1"}); err == nil {
		t.Fatal("bad number")
	}
	if err := p.codeLine([]string{"code", "c", "4", "Q1"}); err != nil || len(p.field.Codes) != 1 {
		t.Fatal(err)
	}
}
