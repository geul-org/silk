//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Schema.Validate — width is a power of two ≥ 16; fields have unique names, fit the width, do not overlap and have valid tables
package sig

import "fmt"

// Validate checks the declaration.
func (s Schema) Validate() error {
	if s.Width < 16 || s.Width&(s.Width-1) != 0 {
		return fmt.Errorf("%w: width %d is not a power of two ≥ 16", ErrSchema, s.Width)
	}
	used := make([]uint64, s.Words())
	names := map[string]bool{}
	for _, f := range s.Fields {
		if names[f.Name] || f.Width <= 0 || f.Width > 64 || f.Bit < 0 || f.Bit+f.Width > s.Width {
			return fmt.Errorf("%w: field %q placement", ErrSchema, f.Name)
		}
		m := mask(f.Bit, f.Width, len(used))
		if !orFree(used, m) {
			return fmt.Errorf("%w: field %q overlaps another", ErrSchema, f.Name)
		}
		if err := checkField(f); err != nil {
			return err
		}
		names[f.Name] = true
	}
	return nil
}
