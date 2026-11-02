//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what Index.docPrint — the live documents as identities
package index

func (ix *Index) docPrint() fingerprint {
	var p fingerprint
	for seq := uint32(0); seq < ix.Live.N; seq++ {
		if ix.Live.Get(seq) {
			p.add([]byte(ix.Path(seq)))
		}
	}
	return p
}
