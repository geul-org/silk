//ff:func feature=index type=rule control=sequence
//ff:what checkDoc — a document is one well-formed GEUL document (gogeul core: words, DOC-*, HEAD/BODY placement, END), so its records frame by TAG WORDs
package index

import (
	"fmt"

	"github.com/geul-org/gogeul/pkg/core"
)

// checkDoc runs gogeul's framing; the byte framing of ingest relies on it.
func checkDoc(b []byte) error {
	ws, err := core.ReadWords(b)
	if err != nil {
		return fmt.Errorf("%w: %v", ErrFormat, err)
	}
	if _, err := core.SplitDocument(ws); err != nil {
		return fmt.Errorf("%w: %v", ErrFormat, err)
	}
	return nil
}
