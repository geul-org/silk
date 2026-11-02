//ff:func feature=index type=builder control=sequence
//ff:what Update — add what is new in the archive and build sections declared since
package index

import (
	"context"

	"github.com/geul-org/silk/pkg/archive"
)

func Update(ctx context.Context, ar *archive.Archive, dir string, o Options) (Report, error) {
	return run(ctx, ar, dir, o, false)
}
