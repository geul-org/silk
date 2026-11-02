//ff:func feature=grammar type=helper control=iteration dimension=1
//ff:what Path.String — dotted numbers
package grammar

import (
	"strconv"
	"strings"
)

func (p Path) String() string {
	s := make([]string, len(p))
	for i, n := range p {
		s[i] = strconv.Itoa(n)
	}
	return strings.Join(s, ".")
}
