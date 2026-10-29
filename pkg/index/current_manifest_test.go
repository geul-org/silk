//ff:func feature=index type=reader control=sequence
//ff:what TestCurrentManifest — the manifest CURRENT names; ErrNoIndex in an empty directory
package index

import (
	"errors"
	"testing"
)

func TestCurrentManifest(t *testing.T) {
	d := t.TempDir()
	if _, err := CurrentManifest(d); !errors.Is(err, ErrNoIndex) {
		t.Fatal(err)
	}
	m := Manifest{}
	commit(d, &m)
	commit(d, &m)
	if got, err := CurrentManifest(d); err != nil || got.Generation != 2 {
		t.Fatal(got, err)
	}
}
