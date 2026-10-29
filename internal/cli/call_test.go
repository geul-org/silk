//ff:func feature=cli type=helper control=sequence
//ff:what call — test helper: run a silk command line, returning the exit code, stdout and stderr
package cli

import (
	"bytes"
	"strings"
)

// call splits args on spaces.
func call(args string) (int, string, string) {
	var o, e bytes.Buffer
	code := Run(strings.Fields(args), &o, &e)
	return code, o.String(), e.String()
}
