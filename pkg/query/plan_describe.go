//ff:func feature=query type=helper control=sequence
//ff:what describe — one plan as text
package query

import "fmt"

func (p *Plan) describe() string {
	if p.Pats != nil {
		return fmt.Sprintf("%s: silk, %d patterns", p.Name, len(p.Pats))
	}
	return fmt.Sprintf("%s: %s, %d partitions, %d prefixes, %d fields fixed, covered=%v", p.Name, p.Decl.Kind, len(p.Parts), len(p.Prefix), p.Fixed, p.Covered)
}
