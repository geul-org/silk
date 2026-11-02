//ff:func feature=decl type=builder control=sequence
//ff:what NewSilk — a SILK declaration from a definition text
package decl

import "github.com/geul-org/silk/pkg/sig"

func NewSilk(name, text string) (Decl, error) {
	if _, err := sig.Parse(text); err != nil {
		return Decl{}, err
	}
	return Decl{Name: name, Kind: KindSilk, Silk: text}, nil
}
