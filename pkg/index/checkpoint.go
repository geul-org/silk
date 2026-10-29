//ff:type feature=index type=model
//ff:what checkpoint — a paused merge: section, task kind, input files, output file, last finished partition key, dead rows dropped so far, writer state
package index

import "github.com/park-jun-woo/silk/pkg/seg"

// checkpoint is one entry of compact.json. It is resumed only when the
// section still holds exactly the same inputs.
type checkpoint struct {
	Section string    `json:"section"`
	Kind    string    `json:"kind"`
	Inputs  []string  `json:"inputs"`
	Output  string    `json:"output"`
	Last    []byte    `json:"last"`
	Dropped uint64    `json:"dropped"`
	State   seg.State `json:"state"`
}
