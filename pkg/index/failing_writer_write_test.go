//ff:func feature=index type=helper control=sequence
//ff:what failingWriter.Write — always an error
package index

import "errors"

func (failingWriter) Write([]byte) (int, error) {
	return 0, errors.New("fail")
}
