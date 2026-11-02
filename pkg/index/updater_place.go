//ff:func feature=index type=writer control=sequence
//ff:what updater.place — move a finished temporary segment into seg/ under the next file number
package index

import (
	"os"
	"path/filepath"
)

func (u *updater) place(tmp string, rows uint64) (SegRef, error) {
	fi, err := os.Stat(tmp)
	if err != nil {
		return SegRef{}, err
	}
	name := u.m.nextName(dirSeg, "sseg")
	if err := os.MkdirAll(filepath.Join(u.dir, dirSeg), 0o755); err != nil {
		return SegRef{}, err
	}
	return SegRef{File: name, Rows: rows, Bytes: uint64(fi.Size())}, os.Rename(tmp, filepath.Join(u.dir, name))
}
