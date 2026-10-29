//ff:func feature=sig type=encoder control=selection
//ff:what patternField — add one condition: an enum field must equal the value's code, a one-hot field must have the value's bit
package sig

import "fmt"

// patternField writes into p's Mask and Want (shared backing arrays).
func patternField(p Pattern, f Field, v uint64) error {
	switch f.Kind {
	case Enum:
		c, ok := f.Codes[v]
		if !ok {
			return fmt.Errorf("%w: field %q has no code for %d", ErrValue, f.Name, v)
		}
		putBits(p.Mask, f.Bit, f.Width, ^uint64(0))
		putBits(p.Want, f.Bit, f.Width, c)
		return nil
	case OneHot:
		b, ok := f.Bits[v]
		if !ok {
			return fmt.Errorf("%w: field %q has no bit for %d", ErrValue, f.Name, v)
		}
		putBits(p.Mask, f.Bit+b, 1, 1)
		putBits(p.Want, f.Bit+b, 1, 1)
		return nil
	default:
		return fmt.Errorf("%w: field %q kind %d", ErrUnsupported, f.Name, f.Kind)
	}
}
