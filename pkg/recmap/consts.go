// Package recmap is SILK's sorted record map: for one data pack and one
// Issuer, the byte offset (u32) of every top-level record of the Issuer's
// version documents, sorted by (Kind·Codec, offset), with a key range table
// and a document table. A key is the record's TAG WORD with only the Kind and
// Codec bits kept (core.TargetWord), so the TRIPLE records of a pack are one
// contiguous range. Maps are built from uncompressed packs, are deterministic
// and are read with ReadAt (docs/recmap.md).
package recmap

const (
	magic      = "SILKRMAP"
	version    = 1
	headerSize = 176
	keySize    = 24      // key table entry
	docSize    = 24      // document table entry
	offSize    = 4       // off[] entry
	maxIssuer  = 32      // bytes of the widest Issuer (16 words)
	keySpace   = 1 << 13 // TargetWord values: Kind (bits 12-11) and Codec (bits 6-0)
)
