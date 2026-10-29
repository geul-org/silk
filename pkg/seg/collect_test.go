//ff:func feature=seg type=helper control=sequence
//ff:what collect — test helper: an emit function that records "doc:off" strings
package seg

import "fmt"

// collect returns the emit callback and the slice it fills.
func collect() (func(doc, off uint32) error, *[]string) {
	var out []string
	return func(doc, off uint32) error {
		out = append(out, fmt.Sprintf("%d:%d", doc, off))
		return nil
	}, &out
}
