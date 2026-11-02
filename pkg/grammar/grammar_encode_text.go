//ff:func feature=grammar type=encoder control=iteration dimension=1
//ff:what Grammar.encodeText — `"text"` as UTF8-TEXT, `"text"@tag` as UTF8-LANGUAGE-TEXT in both language forms, `"text"@@tag` with SHORT-CODE only
package grammar

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geul-org/gogeul/pkg/core"
)

func (g *Grammar) encodeText(tok string) ([][]byte, error) {
	end := strings.LastIndex(tok, `"`)
	s, err := strconv.Unquote(tok[:end+1])
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrSyntax, tok)
	}
	lang := tok[end+1:]
	if lang == "" {
		return [][]byte{RecordOf(core.EncodeBytesValue(core.CodecUTF8TextEven, []byte(s)))}, nil
	}
	tag, short := strings.CutPrefix(lang[1:], "@")
	rs, err := g.langRecords(tag, short)
	if err != nil {
		return nil, err
	}
	var out [][]byte
	for _, r := range rs {
		out = append(out, RecordOf(core.EncodeLanguageText(r, []byte(s))))
	}
	return out, nil
}
