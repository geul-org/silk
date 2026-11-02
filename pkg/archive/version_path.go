//ff:func feature=archive type=encoder control=sequence
//ff:what VersionPath — the archive path of a version document
package archive

import "fmt"

func VersionPath(issuer string, id uint64, version uint32) string {
	return fmt.Sprintf("%s/0x%016x/0x%08x.geul", issuer, id, version)
}
