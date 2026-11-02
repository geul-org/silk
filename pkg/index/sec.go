//ff:type feature=index type=model
//ff:what Sec — one opened section
package index

import "github.com/geul-org/silk/pkg/seg"

type Sec struct {
	Section
	Files []*seg.File
}
