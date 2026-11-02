//ff:func feature=grammar type=decoder control=iteration dimension=1
//ff:what tagNames — the name a BIND gives each DECLARE of a body, by target word
package grammar

import "github.com/geul-org/gogeul/pkg/core"

func tagNames(body []core.Record) map[uint16]string {
	names := map[uint64]string{}
	for _, r := range body {
		b, err := core.DecodeBind(r)
		if err != nil {
			continue
		}
		if id, err := core.DecodeLocalRef(b.Subject); err == nil {
			names[id] = string(b.Name)
		}
	}
	out := map[uint16]string{}
	for _, r := range body {
		if d, err := core.DecodeDeclare(r); err == nil && len(d.Issuer) == 0 {
			out[uint16(d.Target)] = names[d.EdgeID]
		}
	}
	return out
}
