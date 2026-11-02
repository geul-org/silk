//ff:type feature=archive type=model
//ff:what Step — one version document step of a pack's inventory: which document, where, or that it was deleted
package archive

import "github.com/geul-org/gogeul/pkg/pack"

type Step struct {
	Issuer  string
	ID      uint64
	Version uint32
	Loc     pack.Loc
	Deleted bool
}
