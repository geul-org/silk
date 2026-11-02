//ff:func feature=seg type=scanner control=iteration dimension=1
//ff:what TestInfoCheck — shapes the footer can hold pass, others fail
package seg

import (
	"errors"
	"testing"
)

func TestInfoCheck(t *testing.T) {
	if err := testInfo().check(); err != nil {
		t.Fatal(err)
	}
	for _, i := range []Info{{RowWidth: 0}, {RowWidth: maxRowSize + 1}, {RowWidth: 8, PartWidth: -1}, {RowWidth: 8, PartWidth: maxRowSize + 1}, {RowWidth: 8, DocOff: -1}, {RowWidth: 8, DocOff: 5}} {
		if err := i.check(); !errors.Is(err, ErrFormat) {
			t.Fatal(i, err)
		}
	}
}
