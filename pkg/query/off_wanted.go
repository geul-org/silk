//ff:func feature=query type=rule control=sequence
//ff:what offWanted — rows without offsets want every record; rows with offsets want theirs
package query

import (
	"slices"

	"github.com/geul-org/silk/pkg/grammar"
)

func offWanted(hits []Hit, rec grammar.Rec) bool {
	return slices.ContainsFunc(hits, func(h Hit) bool { return !h.HasOff || int(h.Off) == rec.Off })
}
