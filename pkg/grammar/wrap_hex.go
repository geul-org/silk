//ff:func feature=grammar type=helper control=sequence
//ff:what wrapHex — a hex error as ErrSyntax
package grammar

import "fmt"

func wrapHex(err error) error {
	if err == nil {
		return nil
	}
	return fmt.Errorf("%w: %v", ErrSyntax, err)
}
