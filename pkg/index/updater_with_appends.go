//ff:func feature=index type=builder control=iteration dimension=1
//ff:what updater.withAppends — the new run merged with a section's append array, dead rows gone; its rows
package index

import (
	"fmt"
	"path/filepath"

	"github.com/geul-org/silk/pkg/seg"
)

func (u *updater) withAppends(s *Section, tmp string) (string, uint64, error) {
	info, err := infoOf(tmp)
	if err != nil {
		return "", 0, err
	}
	out := filepath.Join(u.tmp, fmt.Sprintf("%d.run", u.n.Add(1)))
	ins := []string{tmp}
	for _, a := range s.Appends {
		ins = append(ins, filepath.Join(u.dir, a.File))
	}
	if _, err := mergeKept(ins, info, seg.Hooks{Live: u.live.Get}, out); err != nil {
		return "", 0, err
	}
	rows, err := rowsIn(out)
	return out, rows, err
}
