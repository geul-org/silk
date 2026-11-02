//ff:func feature=index type=writer control=sequence
//ff:what writeAtomic — put bytes at path through path.tmp (synced, then renamed), creating the directory
package index

import (
	"errors"
	"os"
	"path/filepath"
)

func writeAtomic(path string, b []byte) error {
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		return err
	}
	f, err := os.Create(path + ".tmp")
	if err != nil {
		return err
	}
	_, werr := f.Write(b)
	if err := errors.Join(werr, f.Sync(), f.Close()); err != nil {
		return err
	}
	return os.Rename(path+".tmp", path)
}
