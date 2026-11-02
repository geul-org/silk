//ff:func feature=index type=writer control=iteration dimension=1
//ff:what closeSegs — close segments
package index

import "github.com/geul-org/silk/pkg/seg"

func closeSegs(fs []*seg.File) {
	for _, f := range fs {
		f.Close()
	}
}
