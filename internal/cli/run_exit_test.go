//ff:func feature=cli type=scanner control=sequence
//ff:what TestRunExit — exit codes for usage, failure and success
package cli

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

func TestRunExit(t *testing.T) {
	var o, e bytes.Buffer
	if c := Run(context.Background(), nil, strings.NewReader(""), &o, &e); c != exitUsage {
		t.Fatal(c)
	}
	if c := Run(context.Background(), []string{"nope"}, strings.NewReader(""), &o, &e); c != exitUsage {
		t.Fatal(c)
	}
	if c := Run(context.Background(), []string{"drop"}, strings.NewReader(""), &o, &e); c != exitError {
		t.Fatal(c)
	}
	data, ix := project(t)
	if c := Run(context.Background(), append([]string{"list"}, append(ix, data)...), strings.NewReader(""), &o, &e); c != exitOK {
		t.Fatal(c, e.String())
	}
}
