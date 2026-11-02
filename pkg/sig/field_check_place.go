//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Field.checkPlace — mark a field's bits; past the end, a shared bit or a flag wider than 1 is refused
package sig

import "fmt"

func (f *Field) checkPlace(used []bool) error {
	if f.Off+f.Width > len(used) || f.Kind == KindFlag && f.Width != 1 {
		return fmt.Errorf("field %s: bits %d..%d of %d", f.Name, f.Off, f.Off+f.Width, len(used))
	}
	for i := f.Off; i < f.Off+f.Width; i++ {
		if used[i] {
			return fmt.Errorf("field %s shares bit %d", f.Name, i)
		}
		used[i] = true
	}
	return nil
}
