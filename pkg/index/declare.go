//ff:func feature=index type=writer control=iteration dimension=1
//ff:what Declare — record a declaration and its grammar in a new generation, its sections pending
package index

import (
	"encoding/hex"
	"fmt"
	"path/filepath"
	"slices"
	"strings"

	"github.com/geul-org/silk/pkg/decl"
	"github.com/geul-org/silk/pkg/grammar"
)

func Declare(dir string, d decl.Decl, g *grammar.Grammar) error {
	unlock, err := lockDir(dir)
	if err != nil {
		return err
	}
	defer unlock()
	m, err := currentManifest(dir)
	if err != nil {
		return err
	}
	if slices.ContainsFunc(m.Sections, func(s Section) bool { return s.Name == d.Name }) {
		return fmt.Errorf("%w: %s", ErrExists, d.Name)
	}
	infos, err := d.Infos()
	if err != nil {
		return err
	}
	if err := storeGrammar(dir, &m, g); err != nil {
		return err
	}
	h := d.Hash()
	name := hex.EncodeToString(h[:])
	if err := writeAtomic(filepath.Join(dir, dirDecl, name+".json"), d.Encode()); err != nil {
		return err
	}
	for i := range infos {
		m.Sections = append(m.Sections, Section{Name: d.Name, Decl: name, Bundle: i, Pending: true})
	}
	slices.SortStableFunc(m.Sections, func(a, b Section) int { return strings.Compare(a.Name, b.Name) })
	return commit(dir, &m)
}
