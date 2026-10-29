//ff:func feature=index type=reader control=iteration dimension=1
//ff:what packPaths — the data pack file of each PackRef in the archive (`<set dir>/o/<hash>.gpk`)
package index

import (
	"fmt"
	"path/filepath"

	"github.com/geul-org/silk/pkg/archive"
)

// packPaths finds each pack's set by name.
func packPaths(a *archive.Archive, ps []PackRef) ([]string, error) {
	dirs := map[string]string{}
	for _, s := range a.Sets {
		dirs[s.Name] = s.Dir
	}
	out := make([]string, 0, len(ps))
	for _, p := range ps {
		d, ok := dirs[p.Set]
		if !ok {
			return nil, fmt.Errorf("%w: set %q is gone", ErrStale, p.Set)
		}
		out = append(out, filepath.Join(d, "o", p.Hash+".gpk"))
	}
	return out, nil
}
