//ff:func feature=fixture type=encoder control=sequence
//ff:what Label — a LABEL with a BCP 47 language
package fixture

import "github.com/geul-org/gogeul/pkg/core"

func Label(id uint64, subj []byte, lang, text string) core.Record {
	return edge(0x04, id, subj, LangText(lang, text))
}
