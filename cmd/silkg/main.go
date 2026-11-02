//ff:func feature=cli type=command control=sequence
//ff:what main — run silkg until Ctrl-C and exit with its code
package main

import (
	"context"
	"os"
	"os/signal"

	"github.com/geul-org/silk/internal/cli"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	code := cli.Run(ctx, os.Args[1:], os.Stdin, os.Stdout, os.Stderr)
	stop()
	os.Exit(code)
}
