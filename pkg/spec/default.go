//ff:func feature=spec type=builder control=sequence
//ff:what Default — the default spec for Wikidata documents: TRIPLE by Predicate/Object, QUALIFIER by Property/Value, LABEL·ALIAS·DESCRIPTION by language/Subject, SOURCE and REFERENCE-NODE unpartitioned
package spec

// Default indexes every Edge kind of a Wikidata document. Reference keys use
// 6 bytes (TAG WORD and an ID of up to 2 words: Q and P numbers below 2^32),
// language keys 16 bytes; longer values keep a prefix and the exact check
// compares the whole field.
func Default() Spec {
	ref := func(f int) Key { return Key{Path: []int{f}, Width: 6} }
	lang := Key{Path: []int{2, 0}, Width: 16}
	text := func(name string, codec uint8) Tag {
		return Tag{Name: name, Kind: KindEdge, Codec: codec, Partition: lang, Sort: ref(1), MinPartition: 256}
	}
	return Spec{Version: Version, Tags: []Tag{
		{Name: "TRIPLE", Kind: KindEdge, Codec: 0x0A, Partition: ref(1), Sort: ref(3), MinPartition: 256},
		{Name: "QUALIFIER", Kind: KindEdge, Codec: 0x0B, Partition: ref(1), Sort: ref(3), MinPartition: 256},
		text("LABEL", 0x04), text("ALIAS", 0x05), text("DESCRIPTION", 0x06),
		{Name: "SOURCE", Kind: KindEdge, Codec: 0x20, Sort: ref(2)},
		{Name: "REFERENCE-NODE", Kind: KindEdge, Codec: 0x00, Sort: Key{Path: []int{2}, Width: 16}},
	}}
}
