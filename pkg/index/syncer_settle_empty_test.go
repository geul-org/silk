//ff:func feature=index type=helper control=sequence
//ff:what syncer.settleEmpty — settle of a syncer that replayed nothing
package index

func (s *syncer) settleEmpty() error {
	s.sp = newSpill(pathInfo(), s.u.dir, &s.u.n)
	s.u.o.Limit = 1 << 20
	return s.settle()
}
