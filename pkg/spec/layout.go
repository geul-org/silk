//ff:type feature=spec type=model
//ff:what Layout — an Edge codec's leading record fields at fixed positions (field 0 is the Edge Id)
package spec

// Layout lists the record fields of an Edge codec that every record has, in
// order, from the grammar pages (edge/*.md). Raw words after them (such as a
// TRIPLE's Rank) are not fields: reading them needs the codebook.
type Layout struct {
	Name   string
	Fields []FieldClass
}
