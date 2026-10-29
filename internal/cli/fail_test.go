//ff:func feature=cli type=helper control=sequence
//ff:what TestFail — `silk: <error>` and exit code 1
package cli

import (
	"bytes"
	"errors"
	"testing"
)

func TestFail(t *testing.T) {
	var errb bytes.Buffer
	if fail(&errb, errors.New("boom")) != 1 || errb.String() != "silk: boom\n" {
		t.Fatal(errb.String())
	}
}
