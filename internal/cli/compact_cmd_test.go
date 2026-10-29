//ff:func feature=cli type=command control=sequence
//ff:what TestCompactCmd — compaction under the interrupt context; nothing due is exit 0
package cli

import (
	"io"
	"path/filepath"
	"testing"
)

func TestCompactCmd(t *testing.T) {
	data := packedIndex(t)
	if code := compactCmd(filepath.Join(data, ".silk"), indexFlags{threads: 2}, io.Discard, io.Discard); code != 0 {
		t.Fatal(code)
	}
}
