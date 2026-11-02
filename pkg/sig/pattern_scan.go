//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what Pattern.Scan — visit each matching row of a region sorted by bits, jumping over rows that cannot match
package sig

import "github.com/geul-org/silk/pkg/seg"

func (p Pattern) Scan(r seg.Region, fn func(row []byte) error) error {
	key := make([]byte, (p.Bits+7)/8)
	for i := 0; i < r.Len(); {
		next, err := p.step(r, i, key, fn)
		if next < 0 || err != nil {
			return err
		}
		i = next
	}
	return nil
}
