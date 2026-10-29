//ff:func feature=spec type=rule control=sequence
//ff:what CheckKey — a key's width is even and 2..32, its path is a fixed field (or a language sub-record), and a partition key is bounded
package spec

import "fmt"

// CheckKey validates k against layout l; partition keys must not be fields
// of unbounded length (any Value, a whole language-tagged text).
func CheckKey(l Layout, k Key, partition bool) error {
	if len(k.Path) == 0 {
		if k.Width != 0 {
			return fmt.Errorf("%w: %s key has a width but no path", ErrSpec, l.Name)
		}
		return nil
	}
	if k.Width < 2 || k.Width > MaxWidth || k.Width%2 != 0 {
		return fmt.Errorf("%w: %s key width %d (even, 2..%d)", ErrSpec, l.Name, k.Width, MaxWidth)
	}
	if len(k.Path) > 2 || k.Path[0] < 0 || k.Path[0] >= len(l.Fields) {
		return fmt.Errorf("%w: %s field %v is not at a fixed position (%d record fields)", ErrSpec, l.Name, k.Path, len(l.Fields))
	}
	c := l.Fields[k.Path[0]]
	if len(k.Path) == 2 && (c != ClassLangText || k.Path[1] != 0) {
		return fmt.Errorf("%w: %s field %v: reading inside that field needs the codebook", ErrSpec, l.Name, k.Path)
	}
	if partition && len(k.Path) == 1 && c != ClassRef {
		return fmt.Errorf("%w: %s field %d has variable length and cannot be a partition key", ErrSpec, l.Name, k.Path[0])
	}
	return nil
}
