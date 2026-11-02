//ff:func feature=index type=rule control=iteration dimension=1
//ff:what checkSegs — each segment's hash and order
package index

import "github.com/geul-org/silk/pkg/seg"

func checkSegs(fs []*seg.File) error {
	for _, f := range fs {
		if err := f.Verify(); err != nil {
			return err
		}
	}
	return nil
}
