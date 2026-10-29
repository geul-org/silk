//ff:func feature=seg type=rule control=iteration dimension=1
//ff:what File.checkOther — no key of the other partition also has a partition of its own
package seg

import "fmt"

// checkOther looks every other-partition key up in the table.
func (f *File) checkOther() error {
	gs, err := f.otherGroups(nil)
	if err != nil {
		return err
	}
	for _, g := range gs {
		if _, ok := f.Find(g.Key); ok {
			return fmt.Errorf("%w: %s: key %x is both a partition and in the other partition", ErrFormat, f.Path, g.Key)
		}
	}
	return nil
}
