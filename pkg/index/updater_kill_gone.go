//ff:func feature=index type=writer control=iteration dimension=1
//ff:what updater.killGone — documents in packs the archive no longer has die; a rewritten document comes back by its new pack
package index

func (u *updater) killGone() error {
	gone := u.gonePacks()
	if len(gone) == 0 {
		return nil
	}
	t, err := openDocs(u.dir, &u.m)
	if err != nil {
		return err
	}
	defer t.Close()
	for seq := uint32(0); seq < u.m.DocCount(); seq++ {
		if gone[t.At(seq).Pack] {
			u.live.Put(seq, false)
		}
	}
	return nil
}
