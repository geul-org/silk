//ff:func feature=index type=writer control=sequence
//ff:what Drop — remove a declaration's sections in a new generation
package index

import (
	"fmt"
	"slices"
)

func Drop(dir, name string) error {
	unlock, err := lockDir(dir)
	if err != nil {
		return err
	}
	defer unlock()
	m, err := currentManifest(dir)
	if err != nil {
		return err
	}
	kept := slices.DeleteFunc(slices.Clone(m.Sections), func(s Section) bool { return s.Name == name })
	if len(kept) == len(m.Sections) {
		return fmt.Errorf("%w: %s", ErrUnknown, name)
	}
	m.Sections = kept
	if err := commit(dir, &m); err != nil {
		return err
	}
	collect(dir, m)
	return nil
}
