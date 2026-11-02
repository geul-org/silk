//ff:func feature=archive type=reader control=sequence
//ff:what openSet — one set pinned to the manifest gogeul reported, with its object list
package archive

import (
	"context"
	"os"
	"path/filepath"

	"github.com/geul-org/gogeul/pkg/pack"
)

func openSet(ctx context.Context, si pack.SetInfo) (*Set, error) {
	b, err := os.ReadFile(filepath.Join(si.Dir, "m", si.Manifest+".gpm"))
	if err != nil {
		return nil, err
	}
	objs, err := parseObjects(b)
	if err != nil {
		return nil, err
	}
	r, err := pack.Open(ctx, pack.FileStore{Root: si.Dir}, pack.Options{Manifest: si.Manifest})
	if err != nil {
		return nil, err
	}
	return &Set{Name: si.Name, Dir: si.Dir, Manifest: si.Manifest, Generation: si.Generation, Objects: objs, r: r}, nil
}
