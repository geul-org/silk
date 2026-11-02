//ff:func feature=seg type=decoder control=sequence
//ff:what frame — the File over mapped bytes; the footer's counts must add up to the file size
package seg

import "fmt"

func frame(path string, b []byte) (*File, error) {
	info, parts, rows, err := parseFooter(b[len(b)-footerSize:])
	if err != nil {
		return nil, err
	}
	body := uint64(len(b) - footerSize)
	entry := uint64(info.PartWidth) + 16
	if rows > body/uint64(info.RowWidth) || parts > body/entry || rows*uint64(info.RowWidth)+parts*entry != body {
		return nil, fmt.Errorf("%w: %s sizes", ErrFormat, path)
	}
	return &File{Info: info, Rows: rows, Parts: parts, Path: path, b: b, table: b[rows*uint64(info.RowWidth) : body]}, nil
}
