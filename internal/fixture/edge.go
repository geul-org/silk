//ff:func feature=fixture type=encoder control=iteration dimension=1
//ff:what edge — an Edge from its Edge Id and field bytes
package fixture

import "github.com/geul-org/gogeul/pkg/core"

func edge(codec uint8, id uint64, fields ...[]byte) core.Record {
	p := Local(id)
	for _, f := range fields {
		p = append(p, f...)
	}
	ws, _ := core.ReadWords(p)
	return core.Record{Kind: core.KindEdge, Codec: codec, Payload: ws}
}
