//ff:func feature=grammar type=decoder control=sequence
//ff:what Grammar.addBytes — add the body of one document's bytes
package grammar

import "github.com/geul-org/gogeul/pkg/core"

func (g *Grammar) addBytes(b []byte) error {
	ws, _ := core.ReadWords(b)
	doc, err := core.SplitDocument(ws)
	if err != nil {
		return err
	}
	g.addBody(doc.Body)
	return nil
}
