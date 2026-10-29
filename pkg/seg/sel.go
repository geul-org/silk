//ff:type feature=seg type=model
//ff:what Sel — a segment selection: partition key slot (nil = every partition) and sort key slot (nil = every row)
package seg

// Sel picks rows. Part is PartWidth bytes; Sort is SortWidth bytes (a record
// longer than the slot keeps its prefix, so a match is a candidate).
type Sel struct {
	Part []byte
	Sort []byte
}
