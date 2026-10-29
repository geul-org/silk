//ff:func feature=cli type=command control=iteration dimension=1
//ff:what TestRunQuery — results, counts and plans; usage errors (2) and failing conditions, archives, indexes and plans (1)
package cli

import (
	"strings"
	"testing"
)

func TestRunQuery(t *testing.T) {
	data := packedIndex(t)
	if code, o, _ := call("query -triple P106=Q169470 " + data); code != 0 || strings.Count(o, "\n") != 2 {
		t.Fatal(code, o)
	}
	if code, o, _ := call("query -count -label de -match LABEL:1=Q7554 " + data); code != 0 || o != "1 documents, 31 statements\n" {
		t.Fatal(code, o)
	}
	if code, o, _ := call("query -explain -triple P31=Q5 " + data); code != 0 || !strings.Contains(o, "segment lookup") {
		t.Fatal(code, o)
	}
	for args, want := range map[string]int{
		"query " + data:                                    2,
		"query -triple X=1 " + data:                        1,
		"query -triple P31=Q5 " + t.TempDir():              1,
		"query -index /nonexistent -triple P31=Q5 " + data: 1,
		"query -match LABEL:5=Q1 " + data:                  1,
	} {
		if code, _, _ := call(args); code != want {
			t.Fatal(args, code)
		}
	}
}
