//ff:func feature=sig type=builder control=iteration dimension=1
//ff:what plan.emit — every combination of codes from field i on, appended to out
package sig

import "slices"

func (p plan) emit(out [][]byte, bits []byte, i int) [][]byte {
	if i == len(p.fields) {
		return append(out, slices.Clone(bits))
	}
	for _, c := range p.codes[i] {
		put(bits, p.fields[i].Off, p.fields[i].Width, c)
		out = p.emit(out, bits, i+1)
	}
	return out
}
