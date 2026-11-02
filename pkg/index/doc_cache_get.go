//ff:func feature=index type=reader control=sequence
//ff:what docCache.get — a cached document
package index

func (c *docCache) get(seq uint32) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	e, ok := c.by[seq]
	if !ok {
		return nil, false
	}
	c.order.MoveToFront(e)
	return e.Value.(cached).b, true
}
