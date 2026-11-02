//ff:func feature=index type=rule control=iteration dimension=1
//ff:what Index.sameSec — one declaration's bundles against another index's
package index

import "fmt"

func (ix *Index) sameSec(o *Index, name string, ss []*Sec) error {
	for i, s := range ss {
		if a, b := ix.rowPrint(s), o.rowPrint(o.secs[name][i]); a != b {
			return fmt.Errorf("%w: section %s bundle %d", ErrMismatch, name, i)
		}
	}
	return nil
}
