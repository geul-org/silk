//ff:func feature=index type=builder control=sequence
//ff:what spill.finish — flush, then merge the runs down to one; its path, empty when there are no rows
package index

import "github.com/geul-org/silk/pkg/seg"

func (s *spill) finish(fan int) (string, error) {
	if err := s.flush(); err != nil {
		return "", err
	}
	s.buf.Free()
	return mergeDown(s.runs, s.info, fan, seg.Hooks{}, s.name)
}
