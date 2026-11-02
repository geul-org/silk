//ff:type feature=index type=model
//ff:what docCache — a small LRU of decoded documents, in bytes
package index

import (
	"container/list"
	"sync"
)

type docCache struct {
	max, used int64
	order     *list.List
	by        map[uint32]*list.Element
	mu        sync.Mutex
}
