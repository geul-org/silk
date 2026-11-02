//ff:func feature=seg type=builder control=iteration dimension=1
//ff:what Merge — k-way merge sorted segments into w, dropping rows Live refuses and calling Pause now and then; rows dropped
package seg

func Merge(w *Writer, in []*File, hk Hooks) (uint64, error) {
	h := newHeap(in)
	var dropped uint64
	for n := 0; h.Len() > 0; n++ {
		if err := pausePoint(hk, n); err != nil {
			return dropped, err
		}
		c := (*h)[0]
		keep, err := emit(w, c, hk)
		if err != nil {
			return dropped, err
		}
		if !keep {
			dropped++
		}
		advance(h, c)
	}
	return dropped, nil
}
