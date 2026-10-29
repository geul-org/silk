//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what scanPart — one map's TRIPLE range scanned with the byte matcher, each hit checked exactly and placed in its document
package query

import (
	"fmt"

	"github.com/park-jun-woo/silk/pkg/index"
	"github.com/park-jun-woo/silk/pkg/scan"
)

// scanPart adds p's matching documents and Edge Ids to out.
func scanPart(p index.Part, t Triple, o scan.Options, st *Stats, out map[docKey][]uint64) error {
	offs, err := p.Map.Offsets(tripleKey)
	if err != nil {
		return err
	}
	hits, err := scan.Records(p.Data, p.Map.Header.PackSize, offs, scan.Triple{Pred: t.Pred, Obj: t.Obj}, o)
	if err != nil {
		return err
	}
	st.Scanned += len(offs)
	st.Hits += len(hits)
	for _, h := range hits {
		d, ok := p.Map.Doc(h.Off)
		if !ok {
			return fmt.Errorf("%w: offset %d outside every document", ErrVerify, h.Off)
		}
		id, err := verifyTriple(h.Rec, t)
		if err != nil {
			return fmt.Errorf("document 0x%016x offset %d: %w", d.ID, h.Off, err)
		}
		k := docKey{issuer: p.Map.Header.Issuer, doc: d.ID, version: d.Version}
		out[k] = append(out[k], id)
	}
	return nil
}
