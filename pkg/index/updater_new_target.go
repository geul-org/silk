//ff:func feature=index type=builder control=iteration dimension=1
//ff:what updater.newTarget — one declaration's extractor and spills
package index

import (
	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

func (u *updater) newTarget(s *Section, g *grammar.Grammar) (*target, error) {
	d, err := readDecl(u.dir, s.Decl)
	if err != nil {
		return nil, err
	}
	x, err := decl.NewExtractor(d, g)
	if err != nil {
		return nil, err
	}
	infos, err := d.Infos()
	if err != nil {
		return nil, err
	}
	t := &target{x: x, pending: s.Pending}
	for _, i := range infos {
		t.sp = append(t.sp, newSpill(i, u.tmp, &u.n))
	}
	return t, nil
}
