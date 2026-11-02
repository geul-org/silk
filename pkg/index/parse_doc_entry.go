//ff:func feature=index type=decoder control=sequence
//ff:what parseDocEntry — a 40-byte table entry
package index

import "encoding/binary"

func parseDocEntry(b []byte) DocEntry {
	u32, u64 := binary.BigEndian.Uint32, binary.BigEndian.Uint64
	return DocEntry{Pack: u32(b), Issuer: u32(b[4:]), Off: u64(b[8:]), Len: u64(b[16:]), ID: u64(b[24:]), Version: u32(b[32:]), Enc: b[36]}
}
