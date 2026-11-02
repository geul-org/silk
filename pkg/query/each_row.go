//ff:func feature=query type=scanner control=iteration dimension=1
//ff:what eachRow — visit a region's rows
package query

import "github.com/geul-org/silk/pkg/seg"

func eachRow(reg seg.Region, emit func([]byte) error) error {
	for i := 0; i < reg.Len(); i++ {
		if err := emit(reg.Row(i)); err != nil {
			return err
		}
	}
	return nil
}
