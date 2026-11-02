//ff:func feature=query type=builder control=sequence
//ff:what runner.doc — a framed document; counts each document opened once
package query

import "github.com/geul-org/silk/pkg/grammar"

func (r *runner) doc(seq uint32) (*grammar.Doc, error) {
	r.seen[seq] = true
	b, err := r.ix.Bytes(r.ctx, seq)
	if err != nil {
		return nil, err
	}
	return grammar.Frame(b)
}
