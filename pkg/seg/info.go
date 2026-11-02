//ff:type feature=seg type=model
//ff:what Info — a segment's shape: partition key width, row width, where the document number lies in a row, and the declaration hash
package seg

type Info struct {
	PartWidth int
	RowWidth  int
	DocOff    int
	Hash      [32]byte
}
