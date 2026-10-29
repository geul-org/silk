//ff:type feature=sig type=model
//ff:what Field — one signature field: name, kind, bit offset and width, and its value table (enum codes or one-hot bits)
package sig

// Field occupies bits [Bit, Bit+Width) of a signature (bit 0 is the least
// significant bit of word 0). Codes maps a value (for example a QID) to an
// Enum code (1…2^Width-1); Other is the code of values not in Codes (0 keeps
// them unrecorded). Bits maps a value to a OneHot bit (0…Width-1).
type Field struct {
	Name  string
	Kind  FieldKind
	Bit   int
	Width int
	Codes map[uint64]uint64
	Other uint64
	Bits  map[uint64]int
}
