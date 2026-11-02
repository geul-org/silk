//ff:func feature=decl type=encoder control=sequence
//ff:what Decl.Encode — canonical JSON
package decl

import "encoding/json"

func (d Decl) Encode() []byte {
	b, _ := json.Marshal(d)
	return b
}
