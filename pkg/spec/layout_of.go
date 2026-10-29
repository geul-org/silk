//ff:func feature=spec type=rule control=selection
//ff:what LayoutOf — the fixed record fields of the Edge codecs SILK knows (REFERENCE-NODE, LABEL, ALIAS, DESCRIPTION, TRIPLE, QUALIFIER, SOURCE)
//ff:why Stands in for reading codec-spec documents (DECLARE·FIELD); until then only these grammar layouts can carry keys.
package spec

// LayoutOf returns the layout of an Edge codec; any other Kind or codec has
// none, so no key can be declared on it.
func LayoutOf(kind, codec uint8) (Layout, bool) {
	if kind != KindEdge {
		return Layout{}, false
	}
	switch codec {
	case 0x00:
		return Layout{Name: "REFERENCE-NODE", Fields: []FieldClass{ClassRef, ClassRef, ClassAny}}, true
	case 0x04:
		return Layout{Name: "LABEL", Fields: []FieldClass{ClassRef, ClassRef, ClassLangText}}, true
	case 0x05:
		return Layout{Name: "ALIAS", Fields: []FieldClass{ClassRef, ClassRef, ClassLangText}}, true
	case 0x06:
		return Layout{Name: "DESCRIPTION", Fields: []FieldClass{ClassRef, ClassRef, ClassLangText}}, true
	case 0x0A:
		return Layout{Name: "TRIPLE", Fields: []FieldClass{ClassRef, ClassRef, ClassRef, ClassAny}}, true
	case 0x0B:
		return Layout{Name: "QUALIFIER", Fields: []FieldClass{ClassRef, ClassRef, ClassRef, ClassAny}}, true
	case 0x20:
		return Layout{Name: "SOURCE", Fields: []FieldClass{ClassRef, ClassRef, ClassRef}}, true
	}
	return Layout{}, false
}
