//ff:func feature=fixture type=encoder control=sequence
//ff:what Local — the LOCAL Reference record of n
package fixture

import (
	"strconv"

	"github.com/geul-org/silk/pkg/grammar"
)

func Local(n uint64) []byte {
	b, _ := grammar.Ref("#" + strconv.FormatUint(n, 10))
	return b
}
