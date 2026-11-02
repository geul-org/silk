//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestParseRefuses — widths, code ranges, duplicates, reserved codes and every malformed line are refused; nothing is assigned
package sig

import (
	"errors"
	"math"
	"testing"
)

func TestParseRefuses(t *testing.T) {
	head := "silk s 16\nfield a 0 4 TRIPLE P1 enum\n"
	cases := map[string]string{
		"no header":        "field a 0 4 TRIPLE P1 enum\n  code x 4 Q1\n",
		"bad header":       "silk s\n",
		"bad bits":         "silk s x\n",
		"zero bits":        "silk s 0\n",
		"unknown line":     "silk s 8\nnope\n",
		"field short":      "silk s 8\nfield a 0 4\n",
		"field not TRIPLE": "silk s 8\nfield a 0 4 LABEL P1 enum\n",
		"bad width":        "silk s 8\nfield a 0 x TRIPLE P1 enum\n",
		"wide":             "silk s 80\nfield a 0 65 TRIPLE P1 enum\n",
		"negative off":     "silk s 8\nfield a -1 4 TRIPLE P1 enum\n",
		"bad predicate":    "silk s 8\nfield a 0 4 TRIPLE X1 enum\n",
		"bad kind":         "silk s 8\nfield a 0 4 TRIPLE P1 bloom\n",
		"twice":            head + "  code x 4 Q1\nfield a 4 4 TRIPLE P2 enum\n  code y 4 Q2\n",
		"no codes":         "silk s 8\nfield a 0 4 TRIPLE P1 enum\n",
		"past bits":        "silk s 8\nfield a 4 8 TRIPLE P1 enum\n  code x 4 Q1\n",
		"shared bit":       head + "  code x 4 Q1\nfield b 3 4 TRIPLE P2 enum\n  code y 4 Q2\n",
		"flag width":       "silk s 8\nfield a 0 2 TRIPLE P1 flag\n",
		"flag code":        "silk s 8\nfield a 0 1 TRIPLE P1 flag\n  code x 1 Q1\n",
		"code first":       "silk s 8\ncode x 4 Q1\n",
		"code short":       head + "  code x\n",
		"code range":       head + "  code x 16 Q1\n",
		"reserved":         head + "  code x 3 Q1\n",
		"code twice":       head + "  code x 4 Q1\n  code y 4 Q2\n",
		"name twice":       head + "  code x 4 Q1\n  code x 5 Q2\n",
		"value twice":      head + "  code x 4 Q1\n  code y 5 Q1\n",
		"enum no value":    head + "  code x 4\n",
		"enum bad value":   head + "  code x 4 Z\n",
		"enum bad code":    head + "  code x y Q1\n",
		"prefix bits":      "silk s 8\nfield a 0 4 TRIPLE P1 prefix\n  code x 12 Q1\n",
		"prefix long":      "silk s 8\nfield a 0 4 TRIPLE P1 prefix\n  code x 0b10101 Q1\n",
		"prefix zero":      "silk s 8\nfield a 0 4 TRIPLE P1 prefix\n  code x 0b00 Q1\n",
		"prefix dup":       "silk s 8\nfield a 0 4 TRIPLE P1 prefix\n  code x 0b01 Q1\n  code y 0b0100 Q2\n",
		"bucket none":      "silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4\n",
		"bucket range":     "silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4 1900\n",
		"bucket empty":     "silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4 ..\n",
		"bucket order":     "silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4 1900..1800\n",
		"bucket overlap":   "silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4 ..1900\n  code y 5 1800..\n",
		"bucket num":       "silk s 8\nfield a 0 4 TRIPLE P1 bucket\n  code x 4 a..b\n",
		"when bad":         head + "  code x 4 Q1\nwhen a\n",
		"when field":       head + "  code x 4 Q1\nwhen b=x\n",
		"when code":        head + "  code x 4 Q1\nwhen a=y\n",
		"when on two":      head + "  code x 4 Q1\nfield b 4 4 TRIPLE P2 enum\n  code y 4 Q2\nwhen a=x\nwhen b=y\n",
		"when shares":      head + "  code x 4 Q1\nwhen a=x\nfield b 2 4 TRIPLE P2 enum\n  code y 4 Q2\n",
	}
	for name, text := range cases {
		if _, err := Parse(text); !errors.Is(err, ErrDef) {
			t.Errorf("%s: %v", name, err)
		}
	}
	if _, err := Parse("silk s 8\nfield a 0 4 TRIPLE P1 enum\n  code @top 4 Q1\n  code @t 5 Q2\n"); err != nil {
		t.Fatal(err)
	}
	if _, err := parseKind("flag"); err != nil {
		t.Fatal(err)
	}
	if n, err := parseNumber("0x10"); n != 16 || err != nil {
		t.Fatal(n, err)
	}
	if lo, hi, err := parseRange("1800.."); lo != 1800 || hi != math.MaxInt64 || err != nil {
		t.Fatal(lo, hi, err)
	}
}
