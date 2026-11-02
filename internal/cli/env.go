//ff:type feature=cli type=model
//ff:what env — what a command runs in
package cli

import (
	"context"
	"io"
)

type env struct {
	ctx context.Context
	in  io.Reader
	out io.Writer
	err io.Writer
}
