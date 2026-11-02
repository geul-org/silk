//ff:type feature=sig type=model
//ff:what plan — the code lists of one row group: fields and their chosen codes
package sig

type plan struct {
	fields []*Field
	codes  [][]uint64
}
