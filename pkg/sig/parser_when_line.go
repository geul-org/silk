//ff:func feature=sig type=decoder control=sequence
//ff:what parser.whenLine — `when FIELD=CODE`: following fields belong to it; the field is an enum and the code a declared name
package sig

import (
	"errors"
	"fmt"
	"strings"
)

func (p *parser) whenLine(f []string) error {
	name, val, ok := strings.Cut(strings.Join(f[1:], ""), "=")
	if !ok {
		return errors.New("when FIELD=CODE")
	}
	on := p.d.common(name)
	if on == nil || on.Kind != KindEnum {
		return fmt.Errorf("when on %q, which is no common enum field", name)
	}
	e := on.entry(val)
	if e == nil {
		return fmt.Errorf("when: no code %q", val)
	}
	p.field = nil
	p.when = &When{On: on, Entry: e}
	p.d.Whens = append(p.d.Whens, p.when)
	return nil
}
