//ff:func feature=query type=builder control=iteration dimension=1
//ff:what runner.checkDoc — one document's candidates: meaning conditions of a composite section on its statements, then its records at their offsets, or every record of the TAGs for rows without offsets
package query

import "slices"

func (r *runner) checkDoc(p *Plan, hits []Hit) ([]Hit, error) {
	d, err := r.doc(hits[0].Doc)
	if err != nil {
		return nil, err
	}
	if ok, err := r.meaningHolds(p, d); !ok {
		return nil, err
	}
	var out []Hit
	for _, rec := range d.Recs {
		if !slices.Contains(p.Decl.Tags, rec.Tag) || !offWanted(hits, rec) || !r.matches(d, d.At(rec), p.Cond) {
			continue
		}
		out = append(out, Hit{Doc: hits[0].Doc, Off: uint32(rec.Off), HasOff: true, QID: subjectQID(r.ix, d, rec)})
	}
	return out, nil
}
