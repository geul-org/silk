//ff:type feature=sig type=model
//ff:what parser — definition reading state: the open when block and the field `code` lines belong to
package sig

type parser struct {
	d     *Def
	when  *When
	field *Field
}
