//ff:func feature=grammar type=builder control=sequence
//ff:what Grammar.addLang — one registry entry: canonical hex code and BCP 47 tag
package grammar

import (
	"strconv"
	"strings"
)

func (g *Grammar) addLang(code, tag string) {
	if n, err := strconv.ParseUint(strings.TrimPrefix(code, "0x"), 16, 16); err == nil {
		g.Langs[strings.ToLower(tag)] = uint16(n)
	}
}
