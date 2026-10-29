//ff:type feature=spec type=model
//ff:what Tag — one indexed TAG kind: name, Kind and Codec, partition key, sort key, minimum partition size
package spec

// Tag declares how records of one Kind·Codec are indexed. Partitions with
// fewer than MinPartition entries are kept together in the segment's "other"
// partition.
type Tag struct {
	Name         string `json:"name"`
	Kind         uint8  `json:"kind"`
	Codec        uint8  `json:"codec"`
	Partition    Key    `json:"partition"`
	Sort         Key    `json:"sort"`
	MinPartition int    `json:"min_partition"`
}
