//ff:func feature=grammar type=decoder control=iteration dimension=1
//ff:what Frame — cut a document's bytes into top-level records; a record that does not fit is ErrFrame
package grammar

import "fmt"

func Frame(b []byte) (*Doc, error) {
	d := &Doc{B: b, Recs: make([]Rec, 0, len(b)/32)}
	for p := 0; p < len(b); {
		n := RecLen(b[p:])
		if n <= 0 || n > len(b)-p {
			return nil, fmt.Errorf("%w: at byte %d", ErrFrame, p)
		}
		d.Recs = append(d.Recs, Rec{Off: p, Len: n, Tag: Target(b[p:])})
		p += n
	}
	return d, nil
}
