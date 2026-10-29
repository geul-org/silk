//ff:type feature=index type=model
//ff:what Target — one map to build or verify: the set, the data pack, the Issuer and that Issuer's documents in the pack
package index

import "github.com/park-jun-woo/silk/pkg/archive"

// Target names one record map of an archive.
type Target struct {
	Pack   archive.Pack
	Issuer string
	Docs   []archive.Doc
}
