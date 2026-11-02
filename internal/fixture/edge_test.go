//ff:func feature=fixture type=scanner control=sequence
//ff:what TestEdge — an Edge carries its Edge Id first
package fixture

import (
	"testing"

	"github.com/geul-org/gogeul/pkg/core"
)

func TestEdge(t *testing.T) {
	if r := edge(0x0A, 3, Tok("P1")); r.Kind != core.KindEdge || len(r.Payload) != 4 || r.Payload[1] != 3 {
		t.Fatal(r)
	}
}
