//ff:func feature=index type=builder control=sequence
//ff:what syncer.settle — merge the new identity rows, kill every step a later step of its path replaces, and add them to the paths section
package index

import "github.com/geul-org/silk/pkg/seg"

func (s *syncer) settle() error {
	p, err := s.sp.finish(fanIn(s.u.o.Limit))
	if err != nil || p == "" {
		return err
	}
	f, err := seg.Open(p)
	if err != nil {
		return err
	}
	old, err := openSegs(segPaths(s.u.dir, s.u.m.Paths))
	if err != nil {
		f.Close()
		return err
	}
	s.u.killReplaced(f, old)
	closeSegs(old)
	f.Close()
	return s.u.publish(&s.u.m.Paths, p, f.Rows)
}
