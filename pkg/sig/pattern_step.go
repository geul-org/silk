//ff:func feature=sig type=scanner control=sequence
//ff:what Pattern.step — handle the row at i: visit it when it matches, else jump; the next index, -1 at the end
package sig

import (
	"bytes"

	"github.com/geul-org/silk/pkg/seg"
)

func (p Pattern) step(r seg.Region, i int, key []byte, fn func(row []byte) error) (int, error) {
	row := r.Row(i)
	copy(key, row)
	if !p.seek(key) {
		return -1, nil
	}
	if bytes.Equal(key, row[:len(key)]) {
		return i + 1, fn(row)
	}
	return r.Seek(i, key), nil
}
