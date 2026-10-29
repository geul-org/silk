//ff:type feature=seg type=model
//ff:what Group — the rows of one partition key in one segment (its own partition, or its slice of the other partition with the key stripped)
package seg

// Group is a merge input; Key is the partition slot.
type Group struct {
	Key   []byte
	f     *File
	rg    region
	strip int
}
