//ff:func feature=fixture type=encoder control=sequence
//ff:what LangText — a UTF8-LANGUAGE-TEXT record with a BCP 47 language
package fixture

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/silk/pkg/grammar"
)

func LangText(lang, text string) []byte {
	l := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte(lang))
	return grammar.RecordOf(core.EncodeLanguageText(l, []byte(text)))
}
