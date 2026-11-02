//ff:func feature=grammar type=decoder control=iteration dimension=1
//ff:what EntityID — the QID word of a Subject record: Reference codec in the top 4 bits, ID in the low 28; 0 when it is not an identified QID or PID or does not fit
package grammar

import "github.com/geul-org/gogeul/pkg/core"

func EntityID(rec []byte) uint32 {
	if len(rec) < 4 || len(rec) > 6 || rec[1]&7 != 0 {
		return 0
	}
	t := Target(rec)
	c := uint32(t>>3) & 0xF
	if t>>11 != kindReference || c != uint32(core.RefWikidataQID) && c != uint32(core.RefWikidataPID) {
		return 0
	}
	var n uint64
	for _, b := range rec[2:] {
		n = n<<8 | uint64(b)
	}
	if n >= 1<<28 {
		return 0
	}
	return c<<28 | uint32(n)
}
