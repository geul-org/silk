//ff:func feature=index type=reader control=iteration dimension=2
//ff:what closeSegs — close every segment file of every section
package index

import "github.com/park-jun-woo/silk/pkg/seg"

// closeSegs ignores close errors (files are read-only).
func closeSegs(ss [][]*seg.File) {
	for _, fs := range ss {
		for _, f := range fs {
			f.Close()
		}
	}
}
