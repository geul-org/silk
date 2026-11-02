//ff:func feature=index type=builder control=iteration dimension=1
//ff:what mergeDown — merge paths at most fan at a time, level by level, until one is left; inputs are removed
package index

import "github.com/geul-org/silk/pkg/seg"

func mergeDown(paths []string, info seg.Info, fan int, hk seg.Hooks, name func() string) (string, error) {
	for len(paths) > 1 {
		next, err := mergeLevel(paths, info, fan, hk, name)
		if err != nil {
			return "", err
		}
		paths = next
	}
	if len(paths) == 0 {
		return "", nil
	}
	return paths[0], nil
}
