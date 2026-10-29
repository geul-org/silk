//ff:func feature=index type=helper control=sequence
//ff:what TestSpecFile — the stored spec of a manifest
package index

import (
	"path/filepath"
	"testing"
)

func TestSpecFile(t *testing.T) {
	if p := SpecFile("/x", Manifest{Spec: "ab"}); p != filepath.Join("/x", "spec", "ab.json") {
		t.Fatal(p)
	}
}
