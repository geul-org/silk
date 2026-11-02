//ff:func feature=fixture type=scanner control=sequence
//ff:what TestRecords — every builder makes a document core accepts
package fixture

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestRecords(t *testing.T) {
	s := Tok("Q1")
	body := []core.Record{Triple(1, Tok("P31"), s, Tok("Q5"), 0), Label(2, s, "en", "x"), Alias(3, s, "en", "y"), Bind(4, s, "z"), LabelShort(5, s, 0x0784, "w")}
	body = append(body, LabelLocal(6, 50, s, "en", "v")...)
	b, err := core.EncodeDocument(nil, body)
	if err != nil {
		t.Fatal(err)
	}
	ws, _ := core.ReadWords(b)
	if d, err := core.SplitDocument(ws); err != nil || len(d.Body) != len(body) {
		t.Fatal(err)
	}
	defer func() {
		if recover() == nil {
			t.Fatal("bad token must panic")
		}
	}()
	Tok("nope")
}
