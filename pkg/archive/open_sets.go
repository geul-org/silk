//ff:func feature=archive type=reader control=iteration dimension=1
//ff:what openSets — every set of an archive; the first failure stops
package archive

import (
	"context"

	"github.com/geul-org/gogeul/pkg/pack"
)

func openSets(ctx context.Context, infos []pack.SetInfo) ([]*Set, error) {
	var out []*Set
	for _, si := range infos {
		s, err := openSet(ctx, si)
		if err != nil {
			return nil, err
		}
		out = append(out, s)
	}
	return out, nil
}
