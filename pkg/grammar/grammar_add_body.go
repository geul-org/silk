//ff:func feature=grammar type=builder control=iteration dimension=1
//ff:what Grammar.addBody — the fields, names and language registry a document's BODY declares
package grammar

import (
	"github.com/geul-org/gogeul/pkg/codecspec"
	"github.com/geul-org/gogeul/pkg/core"
)

func (g *Grammar) addBody(body []core.Record) {
	s := codecspec.Collect(body)
	for t, fs := range s {
		g.Fields[uint16(t)] = convertFields(fs)
	}
	for t, n := range tagNames(body) {
		g.Names[n] = append(g.Names[n], t)
	}
	for code, tag := range s.Languages() {
		g.addLang(code, tag)
	}
}
