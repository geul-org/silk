//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.scanCond — one condition on one document
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

func (r *runner) scanCond(seq uint32, d *grammar.Doc, c *Cond) ([]Hit, error) {
	if c.Decl != nil && c.Decl.Silk != "" && len(c.Silk) > 0 {
		ok, err := r.scanSilk(seq, d, c)
		if err != nil || !ok {
			return nil, err
		}
	}
	if c.Decl != nil && c.Decl.Kind == decl.KindSilk {
		return []Hit{{Doc: seq}}, nil
	}
	var out []Hit
	for _, rec := range d.Recs {
		if slices.Contains(c.Tags, rec.Tag) && r.matches(d, d.At(rec), c) {
			out = append(out, Hit{Doc: seq, Off: uint32(rec.Off), HasOff: true, QID: subjectQID(r.ix, d, rec)})
		}
	}
	return out, nil
}
