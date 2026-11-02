//ff:func feature=grammar type=encoder control=iteration dimension=1
//ff:what recordsOf — records as byte strings
package grammar

import "github.com/geul-org/gogeul/pkg/core"

func recordsOf(rs []core.Record) [][]byte {
	out := make([][]byte, len(rs))
	for i, r := range rs {
		out[i] = RecordOf(r)
	}
	return out
}
