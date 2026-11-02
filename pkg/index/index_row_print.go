//ff:func feature=index type=scanner control=iteration dimension=1
//ff:what Index.rowPrint — a section's live rows with each document number read as its path
package index

func (ix *Index) rowPrint(s *Sec) fingerprint {
	var p fingerprint
	for _, f := range s.Files {
		for i := uint64(0); i < f.Parts; i++ {
			k, r := f.PartAt(i)
			ix.printRows(&p, f, k, r)
		}
	}
	return p
}
