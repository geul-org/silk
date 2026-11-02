//ff:func feature=decl type=builder control=selection
//ff:what Decl.Infos — the shape of each bundle a declaration fills: SILK has value order then document order, the others one
package decl

import (
	"fmt"

	"github.com/geul-org/silk/pkg/seg"
	"github.com/geul-org/silk/pkg/sig"
)

func (d Decl) Infos() ([]seg.Info, error) {
	h := d.Hash()
	switch d.Kind {
	case KindIndex:
		k := sum(d.Widths)
		return []seg.Info{{PartWidth: d.PartWidth, RowWidth: k + 8, DocOff: k, Hash: h}}, nil
	case KindText:
		return []seg.Info{{PartWidth: d.PartWidth, RowWidth: TextWidth + 8, DocOff: TextWidth, Hash: h}}, nil
	case KindComposite:
		return []seg.Info{{PartWidth: d.PartWidth, RowWidth: TextWidth + 12, DocOff: TextWidth + 8, Hash: h}}, nil
	case KindSilk:
		def, err := sig.Parse(d.Silk)
		if err != nil {
			return nil, err
		}
		n := (def.Bits + 7) / 8
		return []seg.Info{{RowWidth: n + 4, DocOff: n, Hash: h}, {RowWidth: n + 4, DocOff: 0, Hash: h}}, nil
	}
	return nil, fmt.Errorf("%w: kind %q", ErrDecl, d.Kind)
}
