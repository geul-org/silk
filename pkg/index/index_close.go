//ff:func feature=index type=writer control=iteration dimension=1
//ff:what Index.Close — unmap and release the pin
package index

import "os"

func (ix *Index) Close() {
	for _, ss := range ix.secs {
		for _, s := range ss {
			closeSegs(s.Files)
		}
	}
	if ix.Docs != nil {
		ix.Docs.Close()
	}
	os.Remove(ix.pin)
}
