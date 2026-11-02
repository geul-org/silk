//ff:func feature=index type=helper control=sequence
//ff:what withoutSegs — all minus some
package index

import "slices"

func withoutSegs(all, drop []SegRef) []SegRef {
	return slices.DeleteFunc(slices.Clone(all), func(r SegRef) bool { return slices.Contains(drop, r) })
}
