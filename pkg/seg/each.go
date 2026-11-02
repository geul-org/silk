//ff:func feature=seg type=scanner control=iteration dimension=1
//ff:what Each — visit every row of the sorted inputs in merged order, dropping rows Live refuses
package seg

func Each(in []*File, hk Hooks, fn func(part, row []byte) error) error {
	h := newHeap(in)
	for h.Len() > 0 {
		c := (*h)[0]
		if err := visit(c, hk, fn); err != nil {
			return err
		}
		advance(h, c)
	}
	return nil
}
