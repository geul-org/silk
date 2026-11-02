//ff:type feature=sig type=model
//ff:what Pattern — terms in bit order; a row matches when every term holds
package sig

type Pattern struct {
	Bits  int
	Terms []Term
}
