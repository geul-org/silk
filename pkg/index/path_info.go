//ff:func feature=index type=helper control=sequence
//ff:what pathInfo — the shape of identity rows: Issuer number, document ID, version, then document number
package index

import "github.com/geul-org/silk/pkg/seg"

func pathInfo() seg.Info {
	return seg.Info{RowWidth: pathKey + 4, DocOff: pathKey}
}
