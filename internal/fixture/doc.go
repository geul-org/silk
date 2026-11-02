//ff:type feature=fixture type=model
//ff:what Doc — one version document of the test Issuer
package fixture

import "github.com/geul-org/gogeul/pkg/core"

type Doc struct {
	ID      uint64
	Version uint32
	Body    []core.Record
}
