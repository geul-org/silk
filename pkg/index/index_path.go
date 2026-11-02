//ff:func feature=index type=decoder control=sequence
//ff:what Index.Path — the archive path of a document number
package index

import "github.com/geul-org/silk/pkg/archive"

func (ix *Index) Path(seq uint32) string {
	e := ix.Docs.At(seq)
	return archive.VersionPath(ix.Manifest.Issuers[e.Issuer], e.ID, e.Version)
}
