// Package sig is SILK's entity signature: a fixed-width bit string per entity
// (16, 32, 64, 128… bits, declared per index) whose fields a schema declares —
// enumeration codes (one value), one-hot bits and Bloom bits (several values),
// prefix codes (hierarchies) and range buckets (numbers, times). A query
// becomes a mask and pattern tested over a column of signatures in one fused
// pass. Signatures are an approximate first filter; exact checks follow.
//
// This first step implements enumeration and one-hot fields. Declaring a
// schema as a GEUL document (and pointing an index at its content hash),
// Bloom, prefix and bucket fields and packed 16/32-bit columns are TODO.
package sig

const (
	// Enum stores one value as a code (0 = no value, Other = unlisted).
	Enum FieldKind = iota + 1
	// OneHot sets one bit per listed value present (several values).
	OneHot
	// Bloom hashes values into the field's bits (TODO).
	Bloom
	// Prefix stores a hierarchical code whose prefixes match ancestors (TODO).
	Prefix
	// Bucket stores the range bucket of a number or time (TODO).
	Bucket
)
