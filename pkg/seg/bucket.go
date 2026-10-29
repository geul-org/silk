//ff:type feature=seg type=model
//ff:what Bucket — index rows gathered in memory by partition key, before they are written as one segment
package seg

// Bucket collects rows for one TAG kind. Write turns it into a segment.
type Bucket struct {
	Info  Info
	Min   int
	Rows  uint64
	parts map[string][]byte
}
