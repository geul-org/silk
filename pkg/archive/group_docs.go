//ff:func feature=archive type=builder control=iteration dimension=1
//ff:what groupDocs — the live version documents of a state grouped by data pack object number; compressed ones are ErrCompressed
package archive

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/pack"
)

// groupDocs keeps the version paths of st (not versions.geul, h/ or named
// paths) and groups them by object; n is the number of manifest objects.
func groupDocs(st map[string]pack.Loc, n int) (map[uint32][]Doc, error) {
	by := map[uint32][]Doc{}
	for p, l := range st {
		d, ok := parseVersionPath(p)
		if !ok {
			continue
		}
		if l.Enc != encRaw {
			return nil, fmt.Errorf("%w: %s", ErrCompressed, p)
		}
		if int(l.Obj) >= n {
			return nil, fmt.Errorf("%w: %s lies in object %d of %d", ErrManifest, p, l.Obj, n)
		}
		d.Off, d.Len = l.Off, l.Len
		by[l.Obj] = append(by[l.Obj], d)
	}
	return by, nil
}
