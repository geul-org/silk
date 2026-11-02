//ff:func feature=fixture type=encoder control=sequence
//ff:what Bind — a BIND of a name to a Subject
package fixture

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/pkg/grammar"
)

func Bind(id uint64, subj []byte, name string) core.Record {
	return edge(core.EdgeBind, id, subj, grammar.RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(name))))
}
