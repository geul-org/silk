//ff:func feature=scan type=helper control=sequence
//ff:what errReader.ReadAt — always an error
package scan

import "errors"

// ReadAt fails.
func (errReader) ReadAt([]byte, int64) (int, error) {
	return 0, errors.New("read failure")
}
