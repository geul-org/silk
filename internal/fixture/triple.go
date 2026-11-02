//ff:func feature=fixture type=encoder control=sequence
//ff:what Triple — a TRIPLE; rank 0 normal, 1 preferred, 2 deprecated
package fixture

import (
	"encoding/binary"

	"github.com/geul-org/gogeul/pkg/core"
)

func Triple(id uint64, pred, subj, obj []byte, rank uint16) core.Record {
	return edge(core.EdgeTriple, id, pred, subj, obj, binary.BigEndian.AppendUint16(nil, rank))
}
