//ff:func feature=sig type=rule control=sequence
//ff:what Entity.stated — the entity has this value token on the field's predicate, listed in the definition or not
package sig

import (
	"bytes"
	"slices"

	"github.com/geul-org/silk/pkg/grammar"
)

func (e *Entity) stated(f *Field, v string) bool {
	b, err := grammar.Ref(v)
	if err != nil {
		return false
	}
	return slices.ContainsFunc(e.Objects[string(f.Pred)], func(o []byte) bool { return bytes.Equal(o, b) })
}
