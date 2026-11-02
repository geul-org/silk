//ff:func feature=archive type=reader control=sequence
//ff:what TestOpenSets — sets open in turn; a set that cannot be read stops
package archive

import (
	"context"
	"testing"

	"github.com/geul-org/gogeul/pkg/pack"
)

func TestOpenSets(t *testing.T) {
	_, a := testArchive(t)
	s := a.Sets[0]
	si := pack.SetInfo{Name: s.Name, Dir: s.Dir, Manifest: s.Manifest}
	if out, err := openSets(context.Background(), []pack.SetInfo{si}); err != nil || len(out) != 1 {
		t.Fatal(err)
	}
	if _, err := openSets(context.Background(), []pack.SetInfo{{Dir: t.TempDir()}}); err == nil {
		t.Fatal("no set")
	}
}
