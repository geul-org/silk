//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestFieldParseCode — each kind reads its form; a flag has none
package sig

import "testing"

func TestFieldParseCode(t *testing.T) {
	for k, c := range map[Kind]string{KindEnum: "4", KindPrefix: "0b1", KindBucket: "4"} {
		rest := []string{"Q1"}
		if k == KindBucket {
			rest = []string{"1..2"}
		}
		f := &Field{Kind: k, Width: 4}
		if e, err := f.parseCode("n", c, rest); err != nil || e.Name != "n" {
			t.Fatal(k, err)
		}
	}
	f := &Field{Kind: KindFlag}
	if _, err := f.parseCode("n", "1", nil); err == nil {
		t.Fatal("flag")
	}
}
