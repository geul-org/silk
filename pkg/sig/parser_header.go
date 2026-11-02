//ff:func feature=sig type=decoder control=sequence
//ff:what parser.header — `silk NAME BITS`
package sig

import (
	"errors"
	"fmt"
	"strconv"
)

func (p *parser) header(f []string) error {
	if len(f) != 3 {
		return errors.New("silk NAME BITS")
	}
	n, err := strconv.Atoi(f[2])
	if err != nil || n <= 0 {
		return fmt.Errorf("bits %q", f[2])
	}
	p.d.Name, p.d.Bits = f[1], n
	return nil
}
