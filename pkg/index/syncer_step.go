//ff:func feature=index type=builder control=sequence
//ff:what syncer.step — one version step: a table entry, an identity row, a live bit
package index

import (
	"encoding/binary"

	"github.com/geul-org/silk/pkg/archive"
)

func (s *syncer) step(st archive.Step) error {
	seq := s.next
	s.next++
	e := DocEntry{Pack: s.pack, Issuer: s.issuer(st.Issuer), Off: st.Loc.Off, Len: st.Loc.Len, ID: st.ID, Version: st.Version, Enc: st.Loc.Enc}
	if st.Deleted {
		e.Enc = encDead
	}
	if err := s.dw.add(e); err != nil {
		return err
	}
	s.sp.buf.Add(nil, binary.BigEndian.AppendUint32(identity(e), seq))
	s.u.live.Put(seq, !st.Deleted)
	if s.sp.buf.Size() < s.u.o.Limit/2 {
		return nil
	}
	return s.sp.flush()
}
