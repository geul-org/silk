//ff:func feature=sig type=decoder control=sequence
//ff:what parser.codeLine — `code NAME …` under the open field, read by its kind and checked against the field
package sig

import "errors"

func (p *parser) codeLine(f []string) error {
	if p.field == nil || len(f) < 3 {
		return errors.New("code NAME NUMBER … under a field")
	}
	e, err := p.field.parseCode(f[1], f[2], f[3:])
	if err != nil {
		return err
	}
	return p.field.add(e)
}
