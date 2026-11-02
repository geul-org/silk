//ff:func feature=seg type=helper control=sequence
//ff:what testInfo — partitions of 2 bytes, rows of 6 with the document number last
package seg

func testInfo() Info {
	return Info{PartWidth: 2, RowWidth: 6, DocOff: 2, Hash: [32]byte{7}}
}
