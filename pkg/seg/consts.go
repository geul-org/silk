// Package seg is SILK's segment file: the index entries of one TAG kind,
// grouped by partition key. A row is the sort key slot, then the document
// sequence number and the record's offset in the document (big-endian u32
// each), so comparing rows as bytes orders them by (sort key, document,
// offset). Partitions with fewer rows than the minimum go to the "other"
// partition, whose rows carry their partition slot in front. A segment is
// either a sorted run (binary search) or an appended segment (scan). Files
// are written once, through a temporary name, and never changed
// (docs/segments.md).
package seg

const (
	magic      = "SILKSEG1"
	version    = 1
	footerSize = 128
	flagSorted = 1
	seqSize    = 8    // document sequence number and record offset
	tableExtra = 16   // partition table entry: key, then start and count (u64 each)
	chunkRows  = 4096 // rows per read when scanning or merging
)
