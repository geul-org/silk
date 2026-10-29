//ff:func feature=index type=helper control=sequence
//ff:what mapRel — a target's map file relative to the index directory: `<Issuer>/<pack SHA-256>.srm`
package index

// mapRel is slash-separated, as the catalog lists it.
func mapRel(t Target) string {
	return t.Issuer + "/" + t.Pack.Hash + mapExt
}
