//ff:func feature=cli type=command control=selection
//ff:what runIndex — `silk index build|update|add|compact|verify|stat|spec`
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
	if args[0] == "spec" {
		return indexSpec(stdout)
	}
	var f indexFlags
	fs := newIndexFlags(args[0], stderr, &f)
	if fs.Parse(args[1:]) != nil || fs.NArg() != 1 {
		fmt.Fprint(stderr, usageText)
		return 2
	}
	arch, dir := fs.Arg(0), indexDir(f.dir, fs.Arg(0))
	switch args[0] {
	case "build":
		return indexBuild(arch, dir, f, stdout, stderr)
	case "update", "add":
		return indexUpdate(arch, dir, f, stdout, stderr)
	case "compact":
		return compactCmd(dir, f, stdout, stderr)
	case "verify":
		return indexVerify(arch, dir, stdout, stderr)
	case "stat":
		return indexStat(dir, stdout, stderr)
	default:
		fmt.Fprintf(stderr, "silk: unknown index command %q\n\n%s", args[0], usageText)
		return 2
	}
}
