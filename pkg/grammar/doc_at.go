//ff:func feature=grammar type=helper control=sequence
//ff:what Doc.At — the bytes of a top-level record
package grammar

func (d *Doc) At(r Rec) []byte {
	return d.B[r.Off : r.Off+r.Len]
}
