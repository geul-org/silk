//ff:func feature=index type=builder control=sequence
//ff:what Build — rebuild every declared section from scratch
package index

import (
	"context"

	"github.com/geul-org/silk/pkg/archive"
)

func Build(ctx context.Context, ar *archive.Archive, dir string, o Options) (Report, error) {
	return run(ctx, ar, dir, o, true)
}
