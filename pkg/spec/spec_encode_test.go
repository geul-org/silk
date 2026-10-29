//ff:func feature=spec type=encoder control=sequence
//ff:what TestSpecEncode — canonical JSON: deterministic, newline-terminated, round-trips
package spec

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestSpecEncode(t *testing.T) {
	a, b := Default().Encode(), Default().Encode()
	var s Spec
	if !bytes.Equal(a, b) || a[len(a)-1] != '\n' || json.Unmarshal(a, &s) != nil || len(s.Tags) != 7 {
		t.Fatal(string(a))
	}
}
