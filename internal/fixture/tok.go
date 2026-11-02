//ff:func feature=fixture type=encoder control=sequence
//ff:what Tok — the record of a Qn or Pn token; panics on a bad token (test data)
package fixture

import "github.com/geul-org/silk/pkg/grammar"

func Tok(s string) []byte {
	b, err := grammar.Ref(s)
	if err != nil {
		panic(err)
	}
	return b
}
