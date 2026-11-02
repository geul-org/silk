//ff:func feature=index type=writer control=sequence
//ff:what updater.save — write the bitmap, set references, commit and collect
package index

import "path/filepath"

func (u *updater) save() error {
	u.m.Live = u.m.nextName(dirLive, "sbit")
	if err := writeAtomic(filepath.Join(u.dir, u.m.Live), encodeBitmap(u.live)); err != nil {
		return err
	}
	u.m.Sets = setRefs(u.ar)
	if err := commit(u.dir, &u.m); err != nil {
		return err
	}
	collect(u.dir, u.m)
	return nil
}
