//ff:func feature=sig type=decoder control=selection
//ff:what parser.line — dispatch one line by its first word
package sig

import "fmt"

func (p *parser) line(f []string) error {
	if len(f) == 0 {
		return nil
	}
	switch f[0] {
	case "silk":
		return p.header(f)
	case "field":
		return p.fieldLine(f)
	case "when":
		return p.whenLine(f)
	case "code":
		return p.codeLine(f)
	}
	return fmt.Errorf("unknown line %q", f[0])
}
