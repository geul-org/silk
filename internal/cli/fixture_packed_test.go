//ff:func feature=cli type=helper control=sequence
//ff:what fixturePacked — test helper: the packed test archive without an index
package cli

import (
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

// fixturePacked returns the data folder.
func fixturePacked(t testing.TB) string {
	return fixture.Packed(t, nil)
}
