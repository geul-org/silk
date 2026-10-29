//ff:func feature=index type=rule control=sequence
//ff:what TestCheckSeg — a run matching its tag passes; the wrong tag, run flag or a damaged file fails
package index

import (
	"os"
	"testing"

	"github.com/geul-org/silk/pkg/seg"
)

func TestCheckSeg(t *testing.T) {
	a, _, dir := built(t)
	ix, _ := Open(a, dir)
	defer ix.Close()
	f, tg := ix.Segs[0][0], ix.Spec.Tags[0]
	if err := checkSeg(f, tg, true); err != nil {
		t.Fatal(err)
	}
	if checkSeg(f, tg, false) == nil || checkSeg(f, ix.Spec.Tags[1], true) == nil {
		t.Fatal("mismatch")
	}
	b, _ := os.ReadFile(f.Path)
	b[0] ^= 1
	os.WriteFile(f.Path, b, 0o644)
	g, _ := seg.Open(f.Path)
	defer g.Close()
	if checkSeg(g, tg, true) == nil {
		t.Fatal("damaged file")
	}
}
