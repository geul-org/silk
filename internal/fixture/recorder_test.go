//ff:type feature=fixture type=model
//ff:what recorder — a testing.TB that records Fatal without stopping the process
package fixture

import "testing"

type recorder struct {
	testing.TB
	failed bool
}
