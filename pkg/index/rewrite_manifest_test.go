//ff:func feature=index type=helper control=sequence
//ff:what rewriteManifest — test helper: overwrite a generation's manifest file in place with m
package index

import (
	"encoding/json"
	"os"
	"path/filepath"
)

// rewriteManifest changes a committed manifest without a new generation.
func rewriteManifest(dir string, m Manifest) {
	b, _ := json.Marshal(m)
	os.WriteFile(filepath.Join(dir, manifestName(m.Generation)), b, 0o644)
}
