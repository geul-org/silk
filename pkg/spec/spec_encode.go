//ff:func feature=spec type=encoder control=sequence
//ff:what Spec.Encode — the canonical JSON bytes of a spec (fields in declaration order, indented, newline-terminated)
package spec

import "encoding/json"

// Encode is what Hash hashes and what an index stores as spec/<hash>.json.
func (s Spec) Encode() []byte {
	b, _ := json.MarshalIndent(s, "", "  ")
	return append(b, '\n')
}
