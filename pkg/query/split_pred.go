//ff:func feature=query type=decoder control=sequence
//ff:what splitPred — `path=value` or `path^=value`
package query

import "strings"

func splitPred(a string) (string, string, bool, bool) {
	if a == "" || a[0] == '"' {
		return "", "", false, false
	}
	p, v, ok := strings.Cut(a, "=")
	if !ok {
		return "", "", false, false
	}
	path, prefix := strings.CutSuffix(p, "^")
	return path, v, prefix, true
}
