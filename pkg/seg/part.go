//ff:type feature=seg type=model
//ff:what Part — one partition table entry: partition key slot and its rows [Start, Start+Count)
package seg

// Part locates one partition's rows among the partition rows.
type Part struct {
	Key   []byte
	Start uint64
	Count uint64
}
