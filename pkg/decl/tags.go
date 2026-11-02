//ff:func feature=decl type=builder control=sequence
//ff:what Tags — the target words of a TAG name (both members of a byte-string pair) or `0x` word
package decl

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geul-org/silk/pkg/grammar"
)

func Tags(g *grammar.Grammar, tag string) ([]uint16, error) {
	if t, ok := g.Names[tag]; ok {
		return t, nil
	}
	if !strings.HasPrefix(tag, "0x") {
		return nil, fmt.Errorf("%w: no TAG named %q", ErrDecl, tag)
	}
	n, err := strconv.ParseUint(tag[2:], 16, 16)
	if err != nil || len(tag) != 6 {
		return nil, fmt.Errorf("%w: TAG word %q needs four hex digits", ErrDecl, tag)
	}
	if _, ok := g.Fields[uint16(n)]; !ok {
		return nil, fmt.Errorf("%w: no codec spec for %s", ErrDecl, tag)
	}
	return []uint16{uint16(n)}, nil
}
