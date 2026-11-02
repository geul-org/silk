//ff:func feature=sig type=helper control=sequence
//ff:what stripComment — a line without its # comment
package sig

import "strings"

func stripComment(s string) string {
	before, _, _ := strings.Cut(s, "#")
	return before
}
