//ff:type feature=recmap type=model
//ff:what Map — a whole record map in memory: header, key ranges, documents (by offset) and off[] sorted by (key, offset)
package recmap

// Map is what Build makes and Encode writes.
type Map struct {
	Header Header
	Keys   []KeyRange
	Docs   []DocEntry
	Offs   []uint32
}
