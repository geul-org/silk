//ff:func feature=spec type=reader control=sequence
//ff:what Load — a spec from a JSON file (unknown fields rejected), validated
package spec

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

// Load reads and validates the spec at path.
func Load(path string) (Spec, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return Spec{}, err
	}
	d := json.NewDecoder(bytes.NewReader(b))
	d.DisallowUnknownFields()
	var s Spec
	if err := d.Decode(&s); err != nil {
		return Spec{}, fmt.Errorf("%w: %s: %v", ErrSpec, path, err)
	}
	return s, s.Validate()
}
