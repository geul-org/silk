//ff:func feature=sig type=rule control=selection
//ff:what checkField — a field's value table fits its kind and width: enum codes in 1…2^w-1 (Other ≤ max), one-hot bits in 0…w-1
package sig

import "fmt"

// checkField validates Codes, Other and Bits; other kinds are accepted for
// declaration but not encoded yet.
func checkField(f Field) error {
	switch f.Kind {
	case Enum:
		return checkCodes(f)
	case OneHot:
		return checkBits(f)
	case Bloom, Prefix, Bucket:
		return nil
	default:
		return fmt.Errorf("%w: field %q kind %d", ErrSchema, f.Name, f.Kind)
	}
}
