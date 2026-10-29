//ff:func feature=seg type=decoder control=iteration dimension=1
//ff:what parseTable — the partition table: keys strictly ascending, row ranges contiguous from 0, non-empty, covering every partition row
package seg

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// parseTable decodes PartWidth+16-byte entries.
func parseTable(b []byte, i Info) ([]Part, error) {
	w := i.PartWidth + tableExtra
	out := make([]Part, 0, i.Parts)
	next := uint64(0)
	for e := b; len(e) > 0; e = e[w:] {
		p := Part{Key: e[:i.PartWidth:i.PartWidth], Start: binary.BigEndian.Uint64(e[i.PartWidth:]), Count: binary.BigEndian.Uint64(e[i.PartWidth+8:])}
		if p.Start != next || p.Count == 0 || p.Count > i.Rows-next || (len(out) > 0 && bytes.Compare(p.Key, out[len(out)-1].Key) <= 0) {
			return nil, fmt.Errorf("%w: partition table entry %d", ErrFormat, len(out))
		}
		out = append(out, p)
		next += p.Count
	}
	if next != i.Rows {
		return nil, fmt.Errorf("%w: partitions cover %d of %d rows", ErrFormat, next, i.Rows)
	}
	return out, nil
}
