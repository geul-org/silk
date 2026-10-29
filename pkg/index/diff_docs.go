//ff:func feature=index type=rule control=iteration dimension=1
//ff:what diffDocs — compare the archive's live versions with the live document table: same identity and same pack bytes is kept, anything else is new, and indexed versions left over are killed
package index

import "slices"

// diffDocs matches by (Issuer, document, version) and by place: a version
// rewritten into a new pack generation counts as a new version.
func diffDocs(m Manifest, old []DocEntry, live Bitmap, docs []liveDoc) docDiff {
	idx := map[docKey]uint32{}
	for i, e := range old {
		if live.Get(uint32(i)) {
			idx[docKey{m.Issuers[e.Issuer], e.ID, e.Version}] = uint32(i)
		}
	}
	var d docDiff
	for _, ld := range docs {
		k := docKey{ld.Doc.Issuer, ld.Doc.ID, ld.Doc.Version}
		seq, ok := idx[k]
		if ok && sameLoc(m, old[seq], ld) {
			d.kept = append(d.kept, keptDoc{seq: seq, doc: ld})
			delete(idx, k)
			continue
		}
		d.added = append(d.added, ld)
	}
	for _, seq := range idx {
		d.killed = append(d.killed, seq)
	}
	slices.Sort(d.killed)
	return d
}
