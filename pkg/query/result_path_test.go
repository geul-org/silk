//ff:func feature=query type=encoder control=sequence
//ff:what TestResultPath — Issuer, 16-digit document ID, 8-digit version
package query

import "testing"

func TestResultPath(t *testing.T) {
	if p := (Result{Issuer: "0x00010000", Doc: 0x1d52, Version: 1}).Path(); p != "0x00010000/0x0000000000001d52/0x00000001.geul" {
		t.Fatal(p)
	}
}
