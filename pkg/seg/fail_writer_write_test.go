//ff:func feature=seg type=helper control=sequence
//ff:what failWriter.Write — always an error
package seg

import "errors"

func (failWriter) Write([]byte) (int, error) {
	return 0, errors.New("fail")
}
