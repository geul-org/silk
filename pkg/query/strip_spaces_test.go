//ff:func feature=query type=helper control=sequence
//ff:what stripSpaces — test helper: hex with spaces removed
package query

import "strings"

// stripSpaces lets hex constants be grouped by word.
func stripSpaces(s string) string {
	return strings.ReplaceAll(s, " ", "")
}
