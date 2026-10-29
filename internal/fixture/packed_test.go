//ff:func feature=fixture type=builder control=iteration dimension=1
//ff:what TestPacked — raw and zstd sets are written; a failing codec reports through t.Fatal
package fixture

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestPacked(t *testing.T) {
	for _, c := range []pack.Codec{nil, &pack.Zstd{}} {
		out := Packed(t, c)
		if _, err := os.Stat(filepath.Join(out, ".well-known", "geul", "0x00010000", "latest.txt")); err != nil {
			t.Fatal(err)
		}
	}
	tb := &recorder{T: t}
	Packed(tb, badCodec{})
	if !tb.failed {
		t.Fatal("a failing codec must call Fatal")
	}
}
