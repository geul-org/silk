//ff:func feature=query type=decoder control=iteration dimension=1
//ff:what TestParseMatch — a kind and several field conditions; no kind, an unknown kind, a bad field or only wildcards fail
package query

import (
	"errors"
	"testing"
)

func TestParseMatch(t *testing.T) {
	c, err := ParseMatch("LABEL:2.0=@de, 1=Q7554")
	if err != nil || c.Codec != 0x04 || len(c.Fields) != 2 || len(c.Fields[0].Path) != 2 {
		t.Fatal(c, err)
	}
	for _, s := range []string{"LABEL", "NOPE:1=Q5", "LABEL:x=Q5", "LABEL:1=*"} {
		if _, err := ParseMatch(s); !errors.Is(err, ErrSyntax) {
			t.Fatal(s, err)
		}
	}
}
