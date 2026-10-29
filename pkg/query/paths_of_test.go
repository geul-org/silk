//ff:func feature=query type=helper control=iteration dimension=1
//ff:what pathsOf — test helper: "path #e1,#e2" strings of results
package query

import "fmt"

// pathsOf formats results for comparison.
func pathsOf(rs []Result) []string {
	var out []string
	for _, r := range rs {
		out = append(out, fmt.Sprintf("%s %v", r.Path(), r.Edges))
	}
	return out
}
