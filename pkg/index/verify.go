//ff:func feature=index type=rule control=sequence
//ff:what Verify — every segment and table file, then a fresh build's rows against this index's, document numbers read as identities
package index

import (
	"context"
	"os"
	"path/filepath"
	"strconv"

	"github.com/geul-org/silk/pkg/archive"
)

func Verify(ctx context.Context, ar *archive.Archive, dir string, o Options) error {
	ix, err := Open(ctx, ar, dir, 0)
	if err != nil {
		return err
	}
	defer ix.Close()
	if err := ix.checkFiles(); err != nil {
		return err
	}
	fresh := filepath.Join(dir, dirTmp, "verify-"+strconv.Itoa(os.Getpid()))
	defer os.RemoveAll(fresh)
	if err := cloneDecls(dir, fresh, ix.Manifest); err != nil {
		return err
	}
	if _, err := Build(ctx, ar, fresh, o); err != nil {
		return err
	}
	fx, err := Open(ctx, ar, fresh, 0)
	if err != nil {
		return err
	}
	defer fx.Close()
	return ix.same(fx)
}
