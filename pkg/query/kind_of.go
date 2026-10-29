//ff:func feature=query type=helper control=iteration dimension=1
//ff:what kindOf — the Kind and Codec of an Edge layout name (TRIPLE, LABEL, ...)
package query

import "github.com/geul-org/silk/pkg/spec"

// kindOf searches the direct Edge codecs.
func kindOf(name string) (uint8, uint8, bool) {
	for c := 0; c < 0x7C; c++ {
		if l, ok := spec.LayoutOf(spec.KindEdge, uint8(c)); ok && l.Name == name {
			return spec.KindEdge, uint8(c), true
		}
	}
	return 0, 0, false
}
