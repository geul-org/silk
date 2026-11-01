//ff:func feature=cli type=command control=selection
//ff:what Run — dispatch `silkg index …`, `silkg query …` and help; the process exit code (0 ok, 1 failure, 2 usage)
package cli

import (
	"fmt"
	"io"
)

// Run executes one silkg command line (without the program name).
func Run(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usageText)
		return 2
	}
	switch args[0] {
	case "index":
		return runIndex(args[1:], stdout, stderr)
	case "query":
		return runQuery(args[1:], stdout, stderr)
	case "help", "-h", "-help", "--help":
		fmt.Fprint(stdout, usageText)
		return 0
	default:
		fmt.Fprintf(stderr, "silkg: unknown command %q\n\n%s", args[0], usageText)
		return 2
	}
}
