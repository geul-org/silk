//ff:func feature=query type=encoder control=sequence
//ff:what encodeLang — a BCP 47 language tag as the bytes of its LANGUAGE-BCP47-CODE Value record (odd member for an odd byte count)
package query

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// encodeLang is how Wikidata documents write label languages.
func encodeLang(tag string) ([]byte, error) {
	if tag == "" {
		return nil, fmt.Errorf("%w: empty language tag", ErrSyntax)
	}
	ws, _ := core.EncodeRecord(core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(tag)))
	return core.PutWords(ws), nil
}
