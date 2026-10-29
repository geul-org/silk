//ff:func feature=cli type=command control=selection
//ff:what runIndex — `silk index build|verify [-index DIR] ARCHIVE`
package cli

import (
	"fmt"
	"io"
)

// runIndex parses the shared flags and runs the subcommand.
func runIndex(args []string, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usageText)
		return 2
	}
	fs, dir := newFlags("index "+args[0], stderr)
	if fs.Parse(args[1:]) != nil || fs.NArg() != 1 {
		fmt.Fprint(stderr, usageText)
		return 2
	}
	switch args[0] {
	case "build":
		return indexBuild(fs.Arg(0), indexDir(*dir, fs.Arg(0)), stdout, stderr)
	case "verify":
		return indexVerify(fs.Arg(0), indexDir(*dir, fs.Arg(0)), stdout, stderr)
	default:
		fmt.Fprintf(stderr, "silk: unknown index command %q\n\n%s", args[0], usageText)
		return 2
	}
}
