//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Def.checkBits — one bit map for a set of fields read together
package sig

import "fmt"

func (d *Def) checkBits(fs []*Field) error {
	used := make([]bool, d.Bits)
	for _, f := range fs {
		if f.Kind != KindFlag && len(f.Codes) == 0 {
			return fmt.Errorf("field %s has no codes", f.Name)
		}
		if err := f.checkPlace(used); err != nil {
			return err
		}
	}
	return nil
}
