//ff:func feature=sig type=scanner control=sequence
//ff:what TestFieldSpans — flag, bucket and table values
package sig

import (
	"errors"
	"testing"
)

func TestFieldSpans(t *testing.T) {
	d := testDef(t)
	f := d.common("alive")
	if s, err := f.spans("1"); err != nil || s[0].Lo != 1 {
		t.Fatal(s, err)
	}
	if _, err := f.spans("2"); !errors.Is(err, ErrCond) {
		t.Fatal(err)
	}
	b, _ := d.field("born")
	if s, err := b.spans("1900"); err != nil || s[0].Lo != 6 {
		t.Fatal(s, err)
	}
	j, _ := d.field("job")
	if s, err := j.spans("writer"); err != nil || s[0].Lo != 5 {
		t.Fatal(s, err)
	}
}
