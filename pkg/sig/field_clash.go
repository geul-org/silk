//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Field.clash — a repeated name, or a code another entry holds
package sig

import "fmt"

func (f *Field) clash(e *Entry) error {
	for _, o := range f.Codes {
		if o.Name == e.Name || o.clashes(e, f.Width) {
			return fmt.Errorf("field %s: code %s repeats %s", f.Name, e.Name, o.Name)
		}
	}
	return nil
}
