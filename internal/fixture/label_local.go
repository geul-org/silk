//ff:func feature=fixture type=encoder control=sequence
//ff:what LabelLocal — a LABEL whose Object is a LOCAL bound by a REFERENCE-NODE, both records
package fixture

import "github.com/geul-org/gogeul/pkg/core"

func LabelLocal(id, local uint64, subj []byte, lang, text string) []core.Record {
	return []core.Record{edge(0x00, id, Local(local), LangText(lang, text)), edge(0x04, id+1, subj, Local(local))}
}
