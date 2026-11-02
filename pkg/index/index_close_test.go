//ff:func feature=index type=scanner control=sequence
//ff:what TestIndexClose — a half-opened index closes without a table
package index

import (
	"path/filepath"
	"testing"
)

func TestIndexClose(t *testing.T) {
	ix := &Index{pin: filepath.Join(t.TempDir(), "p"), secs: map[string][]*Sec{"a": {{}}}}
	ix.Close()
}
