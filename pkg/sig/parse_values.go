//ff:func feature=sig type=decoder control=iteration dimension=1
//ff:what parseValues — value tokens as records
package sig

import "github.com/geul-org/silk/pkg/grammar"

func parseValues(toks []string) ([][]byte, error) {
	var out [][]byte
	for _, t := range toks {
		b, err := grammar.Ref(t)
		if err != nil {
			return nil, err
		}
		out = append(out, b)
	}
	return out, nil
}
