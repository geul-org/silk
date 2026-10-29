//ff:func feature=cli type=command control=sequence
//ff:what silk CLI entry point — run the command line and exit with its code
package main

import (
	"os"

	"github.com/park-jun-woo/silk/internal/cli"
)

func main() {
	exit(cli.Run(os.Args[1:], os.Stdout, os.Stderr))
}
