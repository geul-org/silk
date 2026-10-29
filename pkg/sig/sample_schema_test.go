//ff:func feature=sig type=helper control=sequence
//ff:what sampleSchema — a 128-bit schema: P31 instance-of enum (4 bits), P106 occupation one-hot (8 bits, straddling words), P21 sex enum
package sig

// sampleSchema is the schema the tests encode with.
func sampleSchema() Schema {
	return Schema{Width: 128, Fields: []Field{
		{Name: "P31", Kind: Enum, Bit: 0, Width: 4, Codes: map[uint64]uint64{5: 1, 515: 2}, Other: 15},
		{Name: "P106", Kind: OneHot, Bit: 60, Width: 8, Bits: map[uint64]int{169470: 0, 82955: 1, 36180: 7}},
		{Name: "P21", Kind: Enum, Bit: 68, Width: 2, Codes: map[uint64]uint64{6581097: 1, 6581072: 2}},
	}}
}
