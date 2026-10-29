//ff:type feature=query type=model
//ff:what Cond — one condition as written: TAG kind (name, Kind, Codec) and the fields it fixes
package query

// Cond holds when a record of Kind·Codec has every field of Fields.
type Cond struct {
	Text   string
	Name   string
	Kind   uint8
	Codec  uint8
	Fields []FieldCond
}
