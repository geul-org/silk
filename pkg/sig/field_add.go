//ff:func feature=sig type=rule control=iteration dimension=1
//ff:what Field.add — take a code: it must fit the width, stay off the reserved codes, and repeat no name, number, range or value
package sig

import "fmt"

func (f *Field) add(e *Entry) error {
	if e.Len > f.Width || e.Len < 64 && e.Code>>e.Len != 0 {
		return fmt.Errorf("field %s: code %s does not fit %d bits", f.Name, e.Name, f.Width)
	}
	if lo, _ := e.span(f.Width); lo < firstCode {
		return fmt.Errorf("field %s: code %s collides with the reserved codes", f.Name, e.Name)
	}
	if err := f.clash(e); err != nil {
		return err
	}
	for _, v := range e.Values {
		if f.byValue[string(v)] != nil {
			return fmt.Errorf("field %s: a value of %s has a code already", f.Name, e.Name)
		}
		f.byValue[string(v)] = e
	}
	f.Codes = append(f.Codes, e)
	return nil
}
