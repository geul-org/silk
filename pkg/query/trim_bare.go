//ff:func feature=query type=builder control=iteration dimension=1
//ff:what trimBare — a document with record hits needs no hit without a record
package query

import "slices"

func trimBare(hs []Hit) []Hit {
	withOff := map[uint32]bool{}
	for _, h := range hs {
		withOff[h.Doc] = withOff[h.Doc] || h.HasOff
	}
	return slices.DeleteFunc(hs, func(h Hit) bool { return !h.HasOff && withOff[h.Doc] })
}
