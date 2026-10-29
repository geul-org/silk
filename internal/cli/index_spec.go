//ff:func feature=cli type=command control=sequence
//ff:what indexSpec — print the default index specification as JSON (a starting point for -spec)
package cli

import (
	"io"

	"github.com/park-jun-woo/silk/pkg/spec"
)

// indexSpec writes spec.Default().
func indexSpec(stdout io.Writer) int {
	stdout.Write(spec.Default().Encode())
	return 0
}
