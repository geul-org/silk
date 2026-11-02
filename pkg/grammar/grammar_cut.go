//ff:func feature=grammar type=decoder control=iteration dimension=1
//ff:what Grammar.cut — skip the fields before i from byte q and cut field i
package grammar

func (g *Grammar) cut(rec []byte, t uint16, fs []Field, i, q int) (Val, bool) {
	tail := tailBytes(fs)
	for j := 0; j < i; j++ {
		n := fieldLen(rec, fs[j], q, tail)
		if n < 0 {
			return Val{}, false
		}
		q += n
	}
	n := fieldLen(rec, fs[i], q, tail)
	if n < 0 {
		return Val{}, false
	}
	v := Val{B: rec[q : q+n], Raw: fs[i].Width != 0, Rest: fs[i].Width == RestWidth}
	if v.Rest && g.Odd[t] && n > 0 {
		v.B = v.B[:n-1]
	}
	return v, true
}
