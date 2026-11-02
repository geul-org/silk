//ff:func feature=cli type=helper control=selection
//ff:what flags.mode — the result unit
package cli

import "github.com/geul-org/silk/pkg/query"

func (f *flags) mode() string {
	switch {
	case f.keys:
		return query.ModeKeys
	case f.records:
		return query.ModeRecords
	}
	return query.ModeDocs
}
