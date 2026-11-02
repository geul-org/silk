//ff:func feature=index type=reader control=sequence
//ff:what updater.loadLive — the live bitmap of the current generation
package index

import "path/filepath"

func (u *updater) loadLive() error {
	if u.m.Live == "" {
		return nil
	}
	b, err := readBitmap(filepath.Join(u.dir, u.m.Live))
	u.live = b
	return err
}
