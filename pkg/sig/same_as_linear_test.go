//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what sameAsLinear — each pattern's jumping scan visits exactly the rows a linear match would
package sig

import (
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func sameAsLinear(t *testing.T, ps []Pattern, rows [][]byte, reg seg.Region) {
	t.Helper()
	for _, p := range ps {
		got := 0
		err := p.Scan(reg, func(r []byte) error {
			if !p.Match(r) {
				t.Fatal("visited a row that does not match")
			}
			got++
			return nil
		})
		if want := countMatches(p, rows); err != nil || got != want {
			t.Fatal(got, want, err)
		}
	}
}
