//ff:func feature=index type=helper control=sequence
//ff:what Index.Path — the archive path of a document sequence number (`<Issuer>/0x<16 hex>/0x<8 hex>.geul`)
package index

import "fmt"

// Path names document doc.
func (ix *Index) Path(doc uint32) string {
	d := ix.Docs[doc]
	return fmt.Sprintf("%s/0x%016x/0x%08x.geul", ix.M.Issuers[d.Issuer], d.ID, d.Version)
}
