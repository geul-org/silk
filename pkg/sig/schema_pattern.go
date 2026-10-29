//ff:func feature=sig type=encoder control=iteration dimension=1
//ff:what Schema.Pattern — the mask and pattern of conditions "field has value" (enum equality, one-hot bit set), ANDed
package sig

import "fmt"

// Pattern turns field name → value conditions into one Pattern.
func (s Schema) Pattern(conds map[string]uint64) (Pattern, error) {
	p := Pattern{Mask: make([]uint64, s.Words()), Want: make([]uint64, s.Words())}
	for name, v := range conds {
		f, ok := s.field(name)
		if !ok {
			return Pattern{}, fmt.Errorf("%w: no field %q", ErrValue, name)
		}
		if err := patternField(p, f, v); err != nil {
			return Pattern{}, err
		}
	}
	return p, nil
}
