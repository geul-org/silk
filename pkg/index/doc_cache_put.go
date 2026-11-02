//ff:func feature=index type=writer control=iteration dimension=1
//ff:what docCache.put — keep a document, dropping the least used past the limit
package index

func (c *docCache) put(seq uint32, b []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	if int64(len(b)) > c.max || c.by[seq] != nil {
		return
	}
	c.by[seq] = c.order.PushFront(cached{seq, b})
	c.used += int64(len(b))
	for c.used > c.max {
		e := c.order.Back()
		v := c.order.Remove(e).(cached)
		delete(c.by, v.seq)
		c.used -= int64(len(v.b))
	}
}
