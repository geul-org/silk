//ff:func feature=index type=builder control=sequence
//ff:what ingest.entry — a document's table entry, registering its data pack and Issuer in the manifest when new
package index

// entry numbers d's pack and Issuer (appending to m.Packs, m.Issuers).
func (g *ingest) entry(d liveDoc, counts []TagCount) DocEntry {
	pk := d.Set + "/" + d.Pack
	p, ok := g.packIx[pk]
	if !ok {
		p = uint32(len(g.m.Packs))
		g.m.Packs = append(g.m.Packs, PackRef{Set: d.Set, Hash: d.Pack})
		g.packIx[pk] = p
	}
	is, ok := g.issIx[d.Doc.Issuer]
	if !ok {
		is = uint32(len(g.m.Issuers))
		g.m.Issuers = append(g.m.Issuers, d.Doc.Issuer)
		g.issIx[d.Doc.Issuer] = is
	}
	return DocEntry{Pack: p, Issuer: is, Off: d.Doc.Off, Len: uint32(d.Doc.Len), ID: d.Doc.ID, Version: d.Doc.Version, Counts: counts}
}
