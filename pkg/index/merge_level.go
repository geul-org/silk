//ff:func feature=index type=builder control=iteration dimension=1
//ff:what mergeLevel — one level of a merge: every fan paths become one
package index

import "github.com/geul-org/silk/pkg/seg"

func mergeLevel(paths []string, info seg.Info, fan int, hk seg.Hooks, name func() string) ([]string, error) {
	var next []string
	for len(paths) > 0 {
		n := min(fan, len(paths))
		p, err := mergeFiles(paths[:n], info, hk, name())
		if err != nil {
			return nil, err
		}
		next = append(next, p)
		paths = paths[n:]
	}
	return next, nil
}
