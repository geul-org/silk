//ff:func feature=index type=helper control=iteration dimension=1
//ff:what eachPack — call eachIssuer for every data pack of a set
package index

import "github.com/park-jun-woo/silk/pkg/archive"

// eachPack visits packs in order.
func eachPack(packs []archive.Pack, fn func(Target) error) error {
	for _, p := range packs {
		if err := eachIssuer(p, fn); err != nil {
			return err
		}
	}
	return nil
}
