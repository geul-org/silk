//ff:func feature=seg type=rule control=iteration dimension=1
//ff:what File.Verify — body SHA-256, row order of each partition (sorted runs), and the other partition grouped by keys absent from the table
package seg

import (
	"crypto/sha256"
	"fmt"
	"io"
)

// Verify checks a segment against its footer.
func (f *File) Verify() error {
	h := sha256.New()
	if _, err := io.Copy(h, io.NewSectionReader(f.r, 0, f.Info.Bytes-footerSize)); err != nil {
		return err
	}
	if [32]byte(h.Sum(nil)) != f.Info.Sum {
		return fmt.Errorf("%w: %s: body SHA-256", ErrFormat, f.Path)
	}
	for i := range f.Parts {
		if err := f.checkOrder(Group{f: f, rg: f.partRegion(i)}, 0, f.Info.Sorted); err != nil {
			return err
		}
	}
	if err := f.checkOrder(Group{f: f, rg: f.otherRegion()}, f.Info.PartWidth, f.Info.Sorted); err != nil {
		return err
	}
	return f.checkOther()
}
