//ff:func feature=seg type=helper control=sequence
//ff:what Info.check — the shape is usable: widths positive and bounded, the document number inside the row
package seg

import "fmt"

func (i Info) check() error {
	if i.RowWidth <= 0 || i.RowWidth > maxRowSize || i.PartWidth < 0 || i.PartWidth > maxRowSize || i.DocOff < 0 || i.DocOff+4 > i.RowWidth {
		return fmt.Errorf("%w: row %d, partition %d, document at %d", ErrFormat, i.RowWidth, i.PartWidth, i.DocOff)
	}
	return nil
}
