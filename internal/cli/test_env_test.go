//ff:func feature=cli type=helper control=sequence
//ff:what testEnv — an env writing to buffers, with this input
package cli

import (
	"bytes"
	"context"
	"io"
	"strings"
)

func testEnv(in string) (*env, *bytes.Buffer) {
	var o bytes.Buffer
	return &env{ctx: context.Background(), in: strings.NewReader(in), out: &o, err: io.Discard}, &o
}
