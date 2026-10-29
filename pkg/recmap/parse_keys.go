//ff:func feature=recmap type=decoder control=iteration dimension=1
//ff:what parseKeys — the key table: keys strictly ascending, ranges contiguous from 0, non-empty, covering all n records
package recmap

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

// parseKeys decodes 24-byte entries and checks they tile off[] exactly.
func parseKeys(b []byte, n uint64) ([]KeyRange, error) {
	le := binary.LittleEndian
	out := make([]KeyRange, 0, len(b)/keySize)
	next := uint64(0)
	for e := b; len(e) > 0; e = e[keySize:] {
		k := KeyRange{Key: le.Uint16(e), Start: le.Uint64(e[8:]), Count: le.Uint64(e[16:])}
		bad := k.Key >= keySpace || !bytes.Equal(e[2:8], make([]byte, 6)) || k.Start != next || k.Count == 0 || k.Count > n-next
		if bad || (len(out) > 0 && k.Key <= out[len(out)-1].Key) {
			return nil, fmt.Errorf("%w: key table entry %d", ErrFormat, len(out))
		}
		out = append(out, k)
		next += k.Count
	}
	if next != n {
		return nil, fmt.Errorf("%w: key ranges cover %d of %d records", ErrFormat, next, n)
	}
	return out, nil
}
