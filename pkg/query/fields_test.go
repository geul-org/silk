//ff:func feature=query type=decoder control=sequence
//ff:what TestFields — splits embedded records by their TAG WORDs; too few records fail
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFields(t *testing.T) {
	ws := []core.Word{0x0088, 1, 0x00a8, 106, 0}
	fs, err := fields(ws, 2)
	if err != nil || len(fs) != 2 || fs[1][1] != 106 {
		t.Fatal(fs, err)
	}
	if _, err := fields(ws[:2], 2); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
}
