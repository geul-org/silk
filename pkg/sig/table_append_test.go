//ff:func feature=sig type=builder control=sequence
//ff:what TestTableAppend — rows of the table's width append; another width fails
package sig

import (
	"errors"
	"testing"
)

func TestTableAppend(t *testing.T) {
	tb := &Table{Words: 2}
	if err := tb.Append([]uint64{1, 2}); err != nil || len(tb.Data) != 2 {
		t.Fatal(err)
	}
	if err := tb.Append([]uint64{1}); !errors.Is(err, ErrValue) {
		t.Fatal(err)
	}
}
