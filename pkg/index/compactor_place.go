//ff:func feature=index type=writer control=sequence
//ff:what compactor.place — move a merged run into seg/
package index

import (
	"os"
	"path/filepath"
)

func (c *compactor) place(tmp string, rows uint64) (SegRef, error) {
	fi, err := os.Stat(tmp)
	if err != nil {
		return SegRef{}, err
	}
	name := c.m.nextName(dirSeg, "sseg")
	return SegRef{File: name, Rows: rows, Bytes: uint64(fi.Size())}, os.Rename(tmp, filepath.Join(c.dir, name))
}
