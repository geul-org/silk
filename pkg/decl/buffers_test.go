//ff:func feature=decl type=builder control=iteration dimension=1
//ff:what buffers — one buffer per bundle of a declaration
package decl

import (
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func buffers(t *testing.T, d Decl) []*seg.Buffer {
	t.Helper()
	is, err := d.Infos()
	if err != nil {
		t.Fatal(err)
	}
	var out []*seg.Buffer
	for _, i := range is {
		out = append(out, seg.NewBuffer(i))
	}
	return out
}
