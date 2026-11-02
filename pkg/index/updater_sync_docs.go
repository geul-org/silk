//ff:func feature=index type=builder control=iteration dimension=1
//ff:what updater.syncDocs — replay the inventory of every data pack not yet seen; give each version step a number, then let later steps kill earlier ones
package index

import (
	"context"
	"path/filepath"
)

func (u *updater) syncDocs(ctx context.Context) error {
	s := &syncer{u: u, next: u.first, iss: map[string]uint32{}}
	for i, n := range u.m.Issuers {
		s.iss[n] = uint32(i)
	}
	var err error
	path := u.m.nextName(dirDoc, "sdoc")
	if s.dw, err = newDocWriter(filepath.Join(u.dir, path), u.first); err != nil {
		return err
	}
	s.sp = newSpill(pathInfo(), u.tmp, &u.n)
	if err := s.replay(ctx); err != nil {
		return err
	}
	if err := s.dw.finish(); err != nil {
		return err
	}
	if s.dw.count > 0 {
		u.m.Docs = append(u.m.Docs, DocFile{File: path, First: u.first, Count: s.dw.count})
	}
	if err := u.killGone(); err != nil {
		return err
	}
	return s.settle()
}
