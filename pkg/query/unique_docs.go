//ff:func feature=query type=helper control=sequence
//ff:what uniqueDocs — one hit per document, first in document order
package query

import (
	"cmp"
	"slices"
)

func uniqueDocs(hs []Hit) []Hit {
	slices.SortStableFunc(hs, func(a, b Hit) int { return cmp.Compare(a.Doc, b.Doc) })
	return slices.CompactFunc(hs, func(a, b Hit) bool { return a.Doc == b.Doc })
}
