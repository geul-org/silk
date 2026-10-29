//ff:func feature=cli type=command control=sequence
//ff:what TestRunQuery — results, counts and plan print; usage, bad conditions, plain trees and missing indexes fail
package cli

import (
	"bytes"
	"strings"
	"testing"

	"github.com/park-jun-woo/silk/internal/fixture"
)

func TestRunQuery(t *testing.T) {
	dir := fixture.Packed(t, nil)
	var out, errb bytes.Buffer
	if runQuery([]string{"-triple", "P106=Q169470", dir}, &out, &errb) != 1 {
		t.Fatal("no index yet")
	}
	Run([]string{"index", "build", dir}, &out, &errb)
	out.Reset()
	if runQuery([]string{"-triple", "P106=Q169470", dir}, &out, &errb) != 0 || strings.Count(out.String(), "\n") != 2 || !strings.Contains(out.String(), "0x5100000000001d82/0x00000001.geul #2") {
		t.Fatal(out.String(), errb.String())
	}
	out.Reset()
	if runQuery([]string{"-count", "-threads", "2", "-triple", "P106=Q169470", "-triple", "P31=Q5", dir}, &out, &errb) != 0 || out.String() != "2 documents, 4 statements\n" {
		t.Fatal(out.String())
	}
	out.Reset()
	if runQuery([]string{"-explain", "-triple", "P106=Q169470", dir}, &out, &errb) != 0 || !strings.Contains(out.String(), "[skip]") {
		t.Fatal(out.String())
	}
	if runQuery([]string{dir}, &out, &errb) != 2 || runQuery([]string{"-triple", "x", dir}, &out, &errb) != 1 {
		t.Fatal("usage, syntax")
	}
	if runQuery([]string{"-triple", "P106=Q169470", fixture.Source()}, &out, &errb) != 1 {
		t.Fatal("plain tree")
	}
}
