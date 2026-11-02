//ff:func feature=index type=builder control=iteration dimension=1
//ff:what compactor.mergeInputs — merge at most a fan-in at a time, the first level keeping the inputs; the result and its rows
package index

import (
	"fmt"
	"path/filepath"

	"github.com/geul-org/silk/pkg/seg"
)

func (c *compactor) mergeInputs(paths []string, info seg.Info, hk seg.Hooks) (string, uint64, error) {
	fan := fanIn(c.o.Limit)
	name := func() string { return filepath.Join(c.tmp, fmt.Sprintf("%d.run", c.n.Add(1))) }
	var next []string
	for len(paths) > 0 {
		n := min(fan, len(paths))
		p, err := mergeKept(paths[:n], info, hk, name())
		if err != nil {
			return "", 0, err
		}
		next = append(next, p)
		paths = paths[n:]
	}
	p, err := mergeDown(next, info, fan, hk, name)
	if err != nil {
		return "", 0, err
	}
	rows, err := rowsIn(p)
	return p, rows, err
}
