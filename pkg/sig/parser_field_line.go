//ff:func feature=sig type=decoder control=sequence
//ff:what parser.fieldLine — `field NAME OFFSET WIDTH TRIPLE PREDICATE KIND`; a field after `when` belongs to that block
package sig

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/geul-org/silk/pkg/grammar"
)

func (p *parser) fieldLine(f []string) error {
	if len(f) != 7 || f[4] != "TRIPLE" {
		return errors.New("field NAME OFFSET WIDTH TRIPLE PREDICATE KIND")
	}
	off, err1 := strconv.Atoi(f[2])
	w, err2 := strconv.Atoi(f[3])
	if err1 != nil || err2 != nil || off < 0 || w <= 0 || w > 64 {
		return fmt.Errorf("offset %q, width %q", f[2], f[3])
	}
	if o, _ := p.d.field(f[1]); o != nil {
		return fmt.Errorf("field %s twice", f[1])
	}
	pred, err := grammar.Ref(f[5])
	if err != nil {
		return err
	}
	kind, err := parseKind(f[6])
	if err != nil {
		return err
	}
	p.field = &Field{Name: f[1], Off: off, Width: w, Pred: pred, Kind: kind, byValue: map[string]*Entry{}}
	p.d.Preds[string(pred)] = true
	if p.when != nil {
		p.when.Fields = append(p.when.Fields, p.field)
		return nil
	}
	p.d.Common = append(p.d.Common, p.field)
	return nil
}
