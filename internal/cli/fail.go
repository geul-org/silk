//ff:func feature=cli type=helper control=sequence
//ff:what fail — print an error as `silkg: …` and return exit code 1
package cli

import (
	"fmt"
	"io"
)

// fail reports err.
func fail(stderr io.Writer, err error) int {
	fmt.Fprintf(stderr, "silkg: %v\n", err)
	return 1
}
