//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what updater.killReplaced — in identity order, every number of a path but the last dies, and so does every older number of that path
package index

import (
	"bytes"

	"github.com/geul-org/silk/pkg/seg"
)

func (u *updater) killReplaced(f *seg.File, old []*seg.File) {
	r := f.All()
	for i := 0; i < r.Len(); i++ {
		row := r.Row(i)
		if i+1 < r.Len() && bytes.Equal(row[:pathKey], r.Row(i + 1)[:pathKey]) {
			u.live.Put(f.Doc(row), false)
		}
		if i == 0 || !bytes.Equal(row[:pathKey], r.Row(i - 1)[:pathKey]) {
			u.killOld(row[:pathKey], old)
		}
	}
}
