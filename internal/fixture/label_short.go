//ff:func feature=fixture type=encoder control=sequence
//ff:what LabelShort — a LABEL with a SHORT-CODE language
package fixture

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/pkg/grammar"
)

func LabelShort(id uint64, subj []byte, code uint16, text string) core.Record {
	l := core.Record{Kind: core.KindValue, Codec: 0x48, Payload: []core.Word{core.Word(code)}}
	return edge(0x04, id, subj, grammar.RecordOf(core.EncodeLanguageText(l, []byte(text))))
}
