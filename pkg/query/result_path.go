//ff:func feature=query type=encoder control=sequence
//ff:what Result.Path — the archive path of the result's version document
package query

import "fmt"

// Path is `<Issuer>/0x<16 hex>/0x<8 hex>.geul`.
func (r Result) Path() string {
	return fmt.Sprintf("%s/0x%016x/0x%08x.geul", r.Issuer, r.Doc, r.Version)
}
