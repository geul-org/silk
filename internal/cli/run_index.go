//ff:func feature=cli type=command control=selection
//ff:what runIndex — the index subcommands
package cli

import "fmt"

func runIndex(e *env, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("index build|update|compact|verify|stat")
	}
	f := newFlags("index", e.err)
	pos, err := f.parseArgs(args[1:])
	if err != nil {
		return err
	}
	if len(pos) != 1 {
		return fmt.Errorf("index %s ARCHIVE", args[0])
	}
	switch args[0] {
	case "build", "update":
		return e.write(f, pos[0], args[0] == "build")
	case "compact":
		return e.compact(f, pos[0])
	case "verify":
		return e.verify(f, pos[0])
	case "stat":
		return e.stat(f, pos[0])
	}
	return fmt.Errorf("index: unknown command %q", args[0])
}
