//ff:func feature=fixture type=encoder control=sequence
//ff:what Alias — an ALIAS with a BCP 47 language
package fixture

import "github.com/geul-org/gogeul/pkg/core"

func Alias(id uint64, subj []byte, lang, text string) core.Record {
	return edge(0x05, id, subj, LangText(lang, text))
}
