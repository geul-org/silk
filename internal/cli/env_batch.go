//ff:func feature=cli type=command control=iteration dimension=1
//ff:what env.batch — one query per input line; each output line starts with the line number
package cli

import (
	"bufio"
	"fmt"
	"strings"

	"github.com/geul-org/silk/pkg/index"
)

func (e *env) batch(ix *index.Index, f *flags) error {
	sc := bufio.NewScanner(e.in)
	for n := 1; sc.Scan(); n++ {
		if err := e.one(ix, f, strings.Fields(sc.Text()), fmt.Sprintf("%d\t", n)); err != nil {
			return err
		}
	}
	return sc.Err()
}
