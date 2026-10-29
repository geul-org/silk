//ff:func feature=index type=rule control=iteration dimension=1
//ff:what rebuildMask — for each spec tag, whether its section must be rebuilt (no section of that name and TAG, or its tag hash changed)
package index

import (
	"encoding/hex"

	"github.com/geul-org/silk/pkg/spec"
)

// rebuildMask compares sp with the manifest's sections.
func rebuildMask(secs []Section, sp spec.Spec) []bool {
	out := make([]bool, len(sp.Tags))
	for i, t := range sp.Tags {
		s := sectionNamed(secs, t.Name)
		h := t.Hash()
		out[i] = s.Tag != spec.TagWord(t.Kind, t.Codec) || s.Hash != hex.EncodeToString(h[:])
	}
	return out
}
