//ff:func feature=index type=builder control=sequence
//ff:what updater.targetOf — a section's target, made once per declaration
package index

import "github.com/geul-org/silk/pkg/grammar"

func (u *updater) targetOf(by map[string]*target, s *Section, g *grammar.Grammar) (*target, error) {
	if t := by[s.Decl]; t != nil {
		return t, nil
	}
	t, err := u.newTarget(s, g)
	by[s.Decl] = t
	return t, err
}
