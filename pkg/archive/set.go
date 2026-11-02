//ff:type feature=archive type=model
//ff:what Set — one pack set at one generation: name, directory, manifest, objects and its reader
package archive

import "github.com/geul-org/gogeul/pkg/pack"

type Set struct {
	Name       string
	Dir        string
	Manifest   string
	Generation uint64
	Objects    []Object
	r          *pack.Reader
}
