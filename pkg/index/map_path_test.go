//ff:func feature=index type=helper control=sequence
//ff:what TestMapPath — a slash-separated catalog entry under the index directory
package index

import (
	"path/filepath"
	"testing"
)

func TestMapPath(t *testing.T) {
	if got := mapPath("/i", "0x0001/ab.srm"); got != filepath.Join("/i", "0x0001", "ab.srm") {
		t.Fatal(got)
	}
}
