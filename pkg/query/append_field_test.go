//ff:func feature=query type=helper control=sequence
//ff:what TestAppendField — a value is added, a wildcard (nil) skipped
package query

import "testing"

func TestAppendField(t *testing.T) {
	fs := appendField(appendField(nil, []int{1}, []byte{1}), []int{2}, nil)
	if len(fs) != 1 || fs[0].Path[0] != 1 {
		t.Fatal(fs)
	}
}
