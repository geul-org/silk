//ff:func feature=cli type=helper control=sequence
//ff:what silkg — run the CLI and give its exit code, output and errors
package cli

import (
	"bytes"
	"context"
	"strings"
)

func silkg(in string, args ...string) (int, string, string) {
	var o, e bytes.Buffer
	c := Run(context.Background(), args, strings.NewReader(in), &o, &e)
	return c, o.String(), e.String()
}
