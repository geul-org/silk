//ff:func feature=cli type=command control=sequence
//ff:what Run — dispatch by the first word; the exit code
package cli

import (
	"context"
	"fmt"
	"io"
)

func Run(ctx context.Context, args []string, stdin io.Reader, stdout, stderr io.Writer) int {
	if len(args) == 0 {
		fmt.Fprint(stderr, usage)
		return exitUsage
	}
	cmds := map[string]func(*env, []string) error{"create": runCreate, "drop": runDrop, "list": runList, "index": runIndex, "query": runQuery}
	fn, ok := cmds[args[0]]
	if !ok {
		fmt.Fprint(stderr, usage)
		return exitUsage
	}
	e := &env{ctx: ctx, in: stdin, out: stdout, err: stderr}
	if err := fn(e, args[1:]); err != nil {
		fmt.Fprintln(stderr, "silkg:", err)
		return exitError
	}
	return exitOK
}
