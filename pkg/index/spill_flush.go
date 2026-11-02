//ff:func feature=index type=writer control=sequence
//ff:what spill.flush — sort the buffer into a run and empty it
package index

import "github.com/geul-org/silk/pkg/seg"

func (s *spill) flush() error {
	if s.buf.Rows() == 0 {
		return nil
	}
	p := s.name()
	w, err := seg.NewWriter(p, s.info)
	if err != nil {
		return err
	}
	if err := s.buf.WriteTo(w); err != nil {
		w.Abort()
		return err
	}
	if _, err := w.Finish(); err != nil {
		return err
	}
	s.runs = append(s.runs, p)
	s.buf.Reset()
	return nil
}
