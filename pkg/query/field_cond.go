//ff:type feature=query type=model
//ff:what FieldCond — one field of a condition: its key path ([field] or [field, 0]) and the record bytes it must equal
package query

// FieldCond requires the field at Path to be exactly Value.
type FieldCond struct {
	Path  []int
	Value []byte
}
