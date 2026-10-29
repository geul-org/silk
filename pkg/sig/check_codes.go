//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what checkCodes — enum codes are nonzero and fit the field width; Other fits too
package sig

import "fmt"

// checkCodes validates an Enum field.
func checkCodes(f Field) error {
	top := uint64(1)<<f.Width - 1
	if f.Width == 64 {
		top = ^uint64(0)
	}
	if f.Other > top {
		return fmt.Errorf("%w: field %q Other code %d", ErrSchema, f.Name, f.Other)
	}
	for v, c := range f.Codes {
		if c == 0 || c > top {
			return fmt.Errorf("%w: field %q code %d for %d", ErrSchema, f.Name, c, v)
		}
	}
	return nil
}
