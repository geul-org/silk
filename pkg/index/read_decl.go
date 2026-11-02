//ff:func feature=index type=reader control=sequence
//ff:what readDecl — a declaration by hash
package index

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/geul-org/silk/pkg/decl"
)

func readDecl(dir, name string) (decl.Decl, error) {
	b, err := os.ReadFile(filepath.Join(dir, dirDecl, name+".json"))
	if err != nil {
		return decl.Decl{}, err
	}
	var d decl.Decl
	if err := json.Unmarshal(b, &d); err != nil {
		return decl.Decl{}, fmt.Errorf("%w: %v", ErrFormat, err)
	}
	return d, nil
}
