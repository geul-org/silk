//ff:func feature=sig type=helper control=sequence
//ff:what testDef — the parsed test definition
package sig

import "testing"

func testDef(t *testing.T) *Def {
	t.Helper()
	d, err := Parse(defText)
	if err != nil {
		t.Fatal(err)
	}
	return d
}
