//ff:func feature=query type=decoder control=sequence
//ff:what TestFields — the first n field records of a TRIPLE payload; a payload too short fails
package query

import (
	"errors"
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestFields(t *testing.T) {
	ws := []core.Word{0x0088, 0x0005, 0x00a8, 0x001f, 0x00a0, 0x002a, 0x00a0, 0x0005, 0x0000}
	fs, err := fields(ws, 4)
	if err != nil || len(fs) != 4 || fs[1][1] != 0x001f {
		t.Fatal(err, fs)
	}
	if _, err := fields(ws[:3], 2); !errors.Is(err, ErrVerify) {
		t.Fatal(err)
	}
}
