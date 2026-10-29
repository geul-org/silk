//ff:func feature=index type=decoder control=sequence
//ff:what recLenAt — the length of the record at p of a checked document (checkDoc guarantees it frames)
package index

import "github.com/park-jun-woo/silk/pkg/spec"

// recLenAt frames one top-level record.
func recLenAt(b []byte, p int) int {
	return spec.RecLen(b[p:])
}
