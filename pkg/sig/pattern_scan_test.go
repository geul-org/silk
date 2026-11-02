//ff:func feature=sig type=scanner control=iteration dimension=1
//ff:what TestPatternScan — jumping visits exactly the rows a linear match would, over every pattern shape
package sig

import (
	"bytes"
	"errors"
	"slices"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestPatternScan(t *testing.T) {
	d := testDef(t)
	var data []byte
	for i := 0; i < 5000; i++ {
		b := make([]byte, 8)
		put(b, 0, 6, uint64(i*7%9))
		put(b, 7, 8, uint64(i*13%11))
		put(b, 15, 6, uint64(i*5%8))
		data = append(append(data, b...), 0, 0, 0, byte(i))
	}
	rows := make([][]byte, 0, 5000)
	for i := 0; i+12 <= len(data); i += 12 {
		rows = append(rows, data[i:i+12])
	}
	slices.SortFunc(rows, bytes.Compare)
	reg := seg.Region{B: bytes.Join(rows, nil), W: 12}
	for _, cs := range [][]Cond{{{"job", "writer"}}, {{"area", "asia"}}, {{"type", "person"}}, {{"born", "old"}}, {{"type", "person"}, {"job", "writer"}, {"area", "korea"}}} {
		ps, err := d.Compile(cs)
		if err != nil {
			t.Fatal(err)
		}
		sameAsLinear(t, ps, rows, reg)
	}
	boom := errors.New("stop")
	ps, _ := d.Compile([]Cond{{"type", "person"}})
	if err := ps[0].Scan(reg, func([]byte) error { return boom }); !errors.Is(err, boom) {
		t.Fatal(err)
	}
}
