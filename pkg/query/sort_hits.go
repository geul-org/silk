//ff:func feature=query type=helper control=sequence
//ff:what sortHits — by document, then offset, one of each
package query

import (
	"bytes"
	"cmp"
	"slices"
)

func sortHits(hs []Hit) []Hit {
	slices.SortFunc(hs, func(a, b Hit) int {
		return cmp.Or(cmp.Compare(a.Doc, b.Doc), cmp.Compare(a.Off, b.Off), bytes.Compare(a.Key, b.Key))
	})
	return slices.CompactFunc(hs, func(a, b Hit) bool {
		return a.Doc == b.Doc && a.Off == b.Off && bytes.Equal(a.Key, b.Key)
	})
}
