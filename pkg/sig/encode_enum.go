//ff:func feature=sig type=encoder control=sequence
//ff:what encodeEnum — an enum field holds at most one value: its code, Other when unlisted, 0 when absent
package sig

import "fmt"

// encodeEnum writes the code of vs[0].
func encodeEnum(s []uint64, f Field, vs []uint64) error {
	if len(vs) > 1 {
		return fmt.Errorf("%w: enum field %q takes one value, got %d", ErrValue, f.Name, len(vs))
	}
	if len(vs) == 0 {
		return nil
	}
	c, ok := f.Codes[vs[0]]
	if !ok {
		c = f.Other
	}
	putBits(s, f.Bit, f.Width, c)
	return nil
}
