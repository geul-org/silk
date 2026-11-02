//ff:func feature=grammar type=reader control=sequence
//ff:what Grammar.addDoc — add one codec spec document's latest version; other documents are skipped
package grammar

import (
	"github.com/geul-org/gogeul/pkg/core"
	"github.com/geul-org/gogeul/pkg/resolve"
)

func (g *Grammar) addDoc(d resolve.Dir, name string) error {
	id, ok := parseDocID(name)
	if !ok {
		return nil
	}
	words, _ := core.ReadWords(id)
	b, _, err := d.Current([]core.Word{0}, words)
	if err != nil || b == nil {
		return err
	}
	return g.addBytes(b)
}
