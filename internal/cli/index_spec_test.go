//ff:func feature=cli type=command control=sequence
//ff:what TestIndexSpec — prints the default spec JSON
package cli

import (
	"strings"
	"testing"
)

func TestIndexSpec(t *testing.T) {
	var b strings.Builder
	if code := indexSpec(&b); code != 0 || !strings.Contains(b.String(), `"min_partition": 256`) {
		t.Fatal(b.String())
	}
}
