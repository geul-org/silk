//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what scanRegion — jump to each pattern or prefix of a region
package query

import "github.com/geul-org/silk/pkg/seg"

func scanRegion(p *Plan, reg seg.Region, emit func([]byte) error) error {
	for _, pat := range p.Pats {
		if err := pat.Scan(reg, emit); err != nil {
			return err
		}
	}
	if p.Pats != nil {
		return nil
	}
	for _, k := range p.Prefix {
		if err := eachRow(reg.Prefix(k), emit); err != nil {
			return err
		}
	}
	return nil
}
