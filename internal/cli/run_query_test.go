//ff:func feature=cli type=scanner control=iteration dimension=1
//ff:what TestRunQuery — one query, a batch, and each way it refuses
package cli

import (
	"strings"
	"testing"
)

func TestRunQuery(t *testing.T) {
	data, f, _ := built(t)
	ix := []string{"-index", f.index}
	e, o := testEnv("TRIPLE predicate=P31 object=Q5\n")
	if err := runQuery(e, append(append([]string{"TRIPLE", "predicate=P31", "object=Q5"}, ix...), data)); err != nil || !strings.Contains(o.String(), ".geul") {
		t.Fatal(err, o)
	}
	if err := runQuery(e, append(append([]string{"-batch"}, ix...), data)); err != nil {
		t.Fatal(err)
	}
	cases := [][]string{{"-nope"}, {}, {"-doc-cache", "x", data}, {"TRIPLE", t.TempDir()}, {"TRIPLE", "-index", t.TempDir(), data}}
	for _, a := range cases {
		if err := runQuery(e, a); err == nil {
			t.Fatal(a)
		}
	}
}
