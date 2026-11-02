//ff:func feature=sig type=scanner control=sequence
//ff:what TestFieldCodesOf — a flag, each marker, a listed and an unlisted value, a prefix pads
package sig

import (
	"testing"

	"github.com/geul-org/silk/internal/fixture"
)

func TestFieldCodesOf(t *testing.T) {
	d := testDef(t)
	job, _ := d.field("job")
	area, _ := d.field("area")
	alive := d.common("alive")
	if c := alive.codesOf([]byte(wordNoValue)); c[0] != 0 {
		t.Fatal(c)
	}
	if c := alive.codesOf([]byte(wordSomeValue)); c[0] != 1 {
		t.Fatal(c)
	}
	if c := job.codesOf([]byte(wordNoValue)); c[0] != CodeNoValue {
		t.Fatal(c)
	}
	if c := job.codesOf([]byte(wordSomeValue)); c[0] != CodeSomeValue {
		t.Fatal(c)
	}
	if c := job.codesOf(fixture.Tok("Q999")); c[0] != CodeOther {
		t.Fatal(c)
	}
	if c := area.codesOf(fixture.Tok("Q884")); c[0] != 16 {
		t.Fatal(c)
	}
	born, _ := d.field("born")
	if c := born.codesOf(edtf("1850")); c[0] != 5 {
		t.Fatal(c)
	}
}
