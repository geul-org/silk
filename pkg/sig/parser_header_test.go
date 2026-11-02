//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestParserHeader — name and bits; bad forms fail
package sig

import "testing"

func TestParserHeader(t *testing.T) {
	p := newParser()
	if err := p.header([]string{"silk", "a", "8"}); err != nil || p.d.Bits != 8 || p.d.Name != "a" {
		t.Fatal(err, p.d)
	}
	for _, f := range [][]string{{"silk"}, {"silk", "a", "x"}, {"silk", "a", "0"}} {
		if err := p.header(f); err == nil {
			t.Fatal(f)
		}
	}
}
