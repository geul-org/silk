//ff:func feature=seg type=helper control=sequence
//ff:what SortGroups — order merge inputs by partition key, keeping input order for equal keys
package seg

import (
	"bytes"
	"slices"
)

// SortGroups sorts gs in place.
func SortGroups(gs []Group) {
	slices.SortStableFunc(gs, func(a, b Group) int { return bytes.Compare(a.Key, b.Key) })
}
