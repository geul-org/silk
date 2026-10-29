// Package spec is SILK's index specification: for each TAG kind (Kind·Codec)
// it declares the partition key and the sort key of the index entries, each
// a record field at a fixed position, and the minimum partition size. A key
// is the leading Width bytes of that field record, zero padded. Positions are
// checked against the record layouts SILK knows (LayoutOf): a field after a
// variable part, a raw word or a position that needs the codebook cannot be a
// key, and a field of unbounded length cannot be a partition key.
package spec

const (
	// Version is the spec format version.
	Version = 1
	// KindEdge is the Kind of Edge records (TAG WORD bits 12-11 = 10).
	KindEdge uint8 = 2
	// MaxWidth is the widest key slot in bytes.
	MaxWidth = 32
	// codecLangTextEven is UTF8-LANGUAGE-TEXT (0x4A, 0x4B odd): a Value whose
	// payload starts with its language record.
	codecLangTextEven uint8 = 0x4A
	kindValue         uint8 = 1
)

// Field classes of a layout.
const (
	// ClassRef is a Reference record (or a one-word slot marker): bounded.
	ClassRef FieldClass = iota + 1
	// ClassAny is a Reference or Value record of any length.
	ClassAny
	// ClassLangText is a UTF8-LANGUAGE-TEXT Value (or a LOCAL Reference); its
	// first embedded record (sub-record 0) is the language record.
	ClassLangText
)
