//ff:func feature=seg type=reader control=sequence
//ff:what Open — a segment file: footer and partition table read and checked against the file size
package seg

import (
	"fmt"
	"os"
)

// Open reads path's footer and table; rows stay on disk.
func Open(path string) (*File, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	sf, err := readTail(f, fi.Size())
	if err != nil {
		f.Close()
		return nil, fmt.Errorf("%s: %w", path, err)
	}
	sf.Path, sf.Info.Bytes, sf.c = path, fi.Size(), f
	return sf, nil
}
