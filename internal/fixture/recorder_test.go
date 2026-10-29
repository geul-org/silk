//ff:type feature=fixture type=model
//ff:what recorder — a testing.TB whose Fatal only records, so a test can see Packed fail
package fixture

import "testing"

// recorder embeds the real T for everything but Fatal.
type recorder struct {
	*testing.T
	failed bool
}
