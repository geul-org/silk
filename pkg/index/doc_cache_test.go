//ff:func feature=index type=scanner control=sequence
//ff:what TestDocCache — hits, least-used eviction, oversized documents and a zero size
package index

import "testing"

func TestDocCache(t *testing.T) {
	c := newDocCache(10)
	c.put(1, make([]byte, 4))
	c.put(2, make([]byte, 4))
	c.get(1)
	c.put(3, make([]byte, 4))
	if _, ok := c.get(2); ok {
		t.Fatal("2 evicted")
	}
	if _, ok := c.get(1); !ok {
		t.Fatal("1 kept")
	}
	c.put(1, make([]byte, 4))
	c.put(9, make([]byte, 11))
	if _, ok := c.get(9); ok {
		t.Fatal("too big")
	}
	if c.used != 8 {
		t.Fatal(c.used)
	}
}
