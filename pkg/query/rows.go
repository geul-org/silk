//ff:func feature=query type=helper control=iteration dimension=1
//ff:what rows — the rows of a section
package query

import "github.com/geul-org/silk/pkg/index"

func rows(s *index.Sec) uint64 {
	var n uint64
	for _, f := range s.Files {
		n += f.Rows
	}
	return n
}
