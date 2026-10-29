//ff:func feature=query type=encoder control=sequence
//ff:what TestResultPath — canonical archive path
package query

import "testing"

func TestResultPath(t *testing.T) {
	if got := (Result{Issuer: "0x00010000", Doc: 0x5100000000001d52, Version: 1}).Path(); got != "0x00010000/0x5100000000001d52/0x00000001.geul" {
		t.Fatal(got)
	}
}
