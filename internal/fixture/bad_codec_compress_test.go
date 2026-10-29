//ff:func feature=fixture type=helper control=sequence
//ff:what badCodec.Compress — always an error
package fixture

import (
	"errors"

	"github.com/geul-org/gogeul/pkg/pack"
)

// Compress fails.
func (badCodec) Compress([]byte, *pack.Dict) ([]byte, error) {
	return nil, errors.New("bad codec")
}
