//ff:func feature=index type=builder control=sequence
//ff:what newDocCache — an empty cache of at most max bytes; 0 keeps nothing
package index

import "container/list"

func newDocCache(max int64) *docCache {
	return &docCache{max: max, order: list.New(), by: map[uint32]*list.Element{}}
}
