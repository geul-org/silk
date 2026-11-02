//ff:type feature=index type=model
//ff:what target — one declaration being filled: its extractor and one spill per bundle
package index

import "github.com/geul-org/silk/pkg/decl"

type target struct {
	sec     []*Section
	x       *decl.Extractor
	sp      []*spill
	pending bool
}
