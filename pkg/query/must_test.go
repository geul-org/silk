//ff:func feature=query type=helper control=sequence
//ff:what must — the words of a successful EncodeRecord (panics on error, test data only)
package query

import "github.com/geul-org/gogeul/pkg/core"

// must unwraps an encoding result.
func must(ws []core.Word, err error) []core.Word {
	if err != nil {
		panic(err)
	}
	return ws
}
