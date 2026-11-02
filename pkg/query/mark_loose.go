//ff:func feature=query type=builder control=iteration dimension=1
//ff:what markLoose — hits of a text key alone stay candidates: the key holds 8 bytes, and the grammar allows U+0000 in text, so even a short key is not the whole text
package query

import "github.com/geul-org/silk/pkg/decl"

func markLoose(p *Plan, hs []Hit) []Hit {
	if p.Decl.Kind == decl.KindIndex || p.Decl.Kind == decl.KindSilk {
		return hs
	}
	for i := range hs {
		hs[i].Loose = true
	}
	return hs
}
