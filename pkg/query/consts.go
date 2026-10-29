// Package query turns SILK conditions into a plan — entity signature filter,
// record map scan, exact check — and runs it over an index. A query is a set
// of TRIPLE conditions (Predicate = Object) that one document must all meet;
// results name the version document and the Edge Ids of the matching TRIPLEs.
package query

import "github.com/geul-org/gogeul/pkg/core"

// tripleKey is the record map key of TRIPLE records: an Edge with Codec 0x0A
// as a TAG WORD with only the Kind and Codec bits set.
var tripleKey = uint16(core.TargetWord(core.KindEdge, core.EdgeTriple))
