//ff:func feature=grammar type=helper control=sequence
//ff:what testLabel — a LABEL on Q42 in a test document: Edge Id 3, text "Hello" in BCP 47 en
package grammar

import "github.com/geul-org/gogeul/pkg/core"

func testLabel() []byte {
	l := core.EncodeBytesValue(core.CodecLanguageBCP47Even, []byte("en"))
	o := RecordOf(core.EncodeLanguageText(l, []byte("Hello")))
	return edgeOf(0x04, tok("#3"), tok("Q42"), o)
}
