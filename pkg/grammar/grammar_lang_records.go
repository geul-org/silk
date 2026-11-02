//ff:func feature=grammar type=encoder control=sequence
//ff:what Grammar.langRecords — the language records of a tag: BCP 47, and SHORT-CODE when the registry has it (one language by grammar); short keeps only SHORT-CODE
package grammar

import (
	"fmt"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

func (g *Grammar) langRecords(tag string, short bool) ([]core.Record, error) {
	code, ok := g.Langs[strings.ToLower(tag)]
	var out []core.Record
	if ok {
		out = append(out, core.Record{Kind: core.KindValue, Codec: 0x48, Payload: []core.Word{core.Word(code)}})
	}
	if !short {
		out = append(out, core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(strings.ToLower(tag))))
	}
	if len(out) == 0 {
		return nil, fmt.Errorf("%w: language %q has no short code", ErrSyntax, tag)
	}
	return out, nil
}
