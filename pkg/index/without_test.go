//ff:func feature=index type=helper control=sequence
//ff:what TestWithout — references whose file is gone are dropped; the result is never nil
package index

import "testing"

func TestWithout(t *testing.T) {
	rs := without([]SegRef{{File: "a"}, {File: "b"}}, map[string]bool{"a": true})
	if len(rs) != 1 || rs[0].File != "b" || without(nil, nil) == nil {
		t.Fatal(rs)
	}
}
