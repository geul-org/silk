//ff:func feature=seg type=rule control=iteration dimension=1
//ff:what File.Verify — the body hash and the order of partitions and rows
package seg

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func (f *File) Verify() error {
	end := uint64(len(f.b) - footerSize)
	s := sha256.Sum256(f.b[:end])
	if !bytes.Equal(s[:], f.b[end+64:end+96]) {
		return fmt.Errorf("%w: %s hash", ErrFormat, f.Path)
	}
	var prev []byte
	var next uint64
	for i := uint64(0); i < f.Parts; i++ {
		k, s, n := f.entry(i)
		if i > 0 && bytes.Compare(k, prev) <= 0 || s != next || n == 0 {
			return fmt.Errorf("%w: %s partition table", ErrFormat, f.Path)
		}
		prev, next = k, s+n
		if err := f.checkRows(i); err != nil {
			return err
		}
	}
	if next != f.Rows {
		return fmt.Errorf("%w: %s row count", ErrFormat, f.Path)
	}
	return nil
}
