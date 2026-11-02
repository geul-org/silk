//ff:type feature=sig type=model
//ff:what When — fields that exist only under one code of a common field
package sig

type When struct {
	On     *Field
	Entry  *Entry
	Fields []*Field
}
