//ff:func feature=recmap type=reader control=sequence
//ff:what Open — a map file from disk (Read over the os.File, which Close releases)
package recmap

import "os"

// Open stats path for its size, opens it and reads its header and tables.
func Open(path string) (*File, error) {
	fi, err := os.Stat(path)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	m, err := Read(f, fi.Size())
	if err != nil {
		f.Close()
		return nil, err
	}
	m.c = f
	return m, nil
}
