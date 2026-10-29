//ff:func feature=cli type=helper control=sequence
//ff:what loadSpec — the -spec file, else the spec the index already has (keep), else the default spec
package cli

import (
	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/spec"
)

// loadSpec picks the spec for build (keep false) or update (keep true).
func loadSpec(path, dir string, keep bool) (spec.Spec, error) {
	if path != "" {
		return spec.Load(path)
	}
	if m, err := index.CurrentManifest(dir); keep && err == nil {
		return spec.Load(index.SpecFile(dir, m))
	}
	return spec.Default(), nil
}
