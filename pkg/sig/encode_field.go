//ff:func feature=sig type=encoder control=selection
//ff:what encodeField — write one field's values by its kind (enum code or one-hot bits); other kinds are ErrUnsupported
package sig

import "fmt"

// encodeField dispatches on f.Kind.
func encodeField(s []uint64, f Field, vs []uint64) error {
	switch f.Kind {
	case Enum:
		return encodeEnum(s, f, vs)
	case OneHot:
		encodeOneHot(s, f, vs)
		return nil
	default:
		return fmt.Errorf("%w: field %q kind %d", ErrUnsupported, f.Name, f.Kind)
	}
}
