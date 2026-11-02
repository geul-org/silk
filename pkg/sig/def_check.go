//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Def.check — fields stay inside the declared bits and never share one, each code-carrying field has codes, when blocks hang on one field
package sig

import (
	"errors"
	"slices"
)

func (d *Def) check() error {
	if err := d.checkBits(d.Common); err != nil {
		return err
	}
	for _, w := range d.Whens {
		if w.On != d.Whens[0].On {
			return errors.New("when blocks on two fields")
		}
		if err := d.checkBits(append(slices.Clone(d.Common), w.Fields...)); err != nil {
			return err
		}
	}
	return nil
}
