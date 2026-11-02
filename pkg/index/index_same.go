//ff:func feature=index type=rule control=iteration dimension=1
//ff:what Index.same — equal live documents and equal live rows per section
package index

import "fmt"

func (ix *Index) same(o *Index) error {
	if a, b := ix.docPrint(), o.docPrint(); a != b {
		return fmt.Errorf("%w: live documents differ", ErrMismatch)
	}
	for name, ss := range ix.secs {
		if err := ix.sameSec(o, name, ss); err != nil {
			return err
		}
	}
	return nil
}
