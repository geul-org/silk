//ff:func feature=seg type=reader control=sequence
//ff:what Open — map a segment read-only and check its footer and sizes
package seg

import (
	"fmt"
)

func Open(path string) (*File, error) {
	b, err := Map(path)
	if err != nil {
		return nil, err
	}
	if len(b) < footerSize {
		Unmap(b)
		return nil, fmt.Errorf("%w: %s too short", ErrFormat, path)
	}
	s, err := frame(path, b)
	if err != nil {
		Unmap(b)
		return nil, err
	}
	return s, nil
}
