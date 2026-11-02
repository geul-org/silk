//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what updater.killOld — clear the numbers older segments hold for one path
package index

import "github.com/geul-org/silk/pkg/seg"

func (u *updater) killOld(key []byte, old []*seg.File) {
	for _, f := range old {
		r := f.All().Prefix(key)
		for i := 0; i < r.Len(); i++ {
			u.live.Put(f.Doc(r.Row(i)), false)
		}
	}
}
