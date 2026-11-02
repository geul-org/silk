//ff:func feature=index type=writer control=iteration dimension=1
//ff:what DocTable.Close — unmap
package index

import "github.com/geul-org/silk/pkg/seg"

func (t *DocTable) Close() {
	for _, f := range t.files {
		seg.Unmap(f.raw)
	}
}
