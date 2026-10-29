//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what checkBits — one-hot bit numbers lie inside the field
package sig

import "fmt"

// checkBits validates a OneHot field.
func checkBits(f Field) error {
	for v, b := range f.Bits {
		if b < 0 || b >= f.Width {
			return fmt.Errorf("%w: field %q bit %d for %d", ErrSchema, f.Name, b, v)
		}
	}
	return nil
}
