//ff:func feature=index type=helper control=sequence
//ff:what syncer.issuer — the number of an Issuer, added when new
package index

func (s *syncer) issuer(name string) uint32 {
	n, ok := s.iss[name]
	if !ok {
		n = uint32(len(s.u.m.Issuers))
		s.iss[name] = n
		s.u.m.Issuers = append(s.u.m.Issuers, name)
	}
	return n
}
