//ff:func feature=cli type=helper control=sequence
//ff:what triples.String — the conditions joined by spaces (flag.Value)
package cli

import "strings"

// String shows the flag's value.
func (t *triples) String() string {
	return strings.Join(*t, " ")
}
