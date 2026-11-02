//ff:func feature=index type=helper control=sequence
//ff:what spill.name — a new temporary run path
package index

import (
	"fmt"
	"path/filepath"
)

func (s *spill) name() string {
	return filepath.Join(s.tmp, fmt.Sprintf("%d.run", s.n.Add(1)))
}
