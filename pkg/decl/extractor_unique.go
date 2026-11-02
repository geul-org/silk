//ff:func feature=decl type=rule control=sequence
//ff:what Extractor.Unique — byte-equal rows of one document mean nothing twice (rows without offsets)
package decl

func (x *Extractor) Unique() bool {
	return x.d.Kind != KindIndex
}
